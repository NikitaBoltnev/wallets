package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/shopspring/decimal"
)

func (h *Handler) PostWalletTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var operation WalletOperationRequest
	var err error
	decoder := json.NewDecoder(r.Body)

	if err = decoder.Decode(&operation); err != nil {
		walletErr := NewErrInvaliJSON()
		statusCode := getHTTPStatusForError(walletErr.Code)
		sendErrorResponse(w, walletErr.Code, walletErr.Message, statusCode)
		return
	}

	if err := operation.validation(); err != nil {
		if walletErr, ok := err.(WalletError); ok {
			statusCode := getHTTPStatusForError(walletErr.Code)
			sendErrorResponse(w, walletErr.Code, walletErr.Message, statusCode)
		}
		return
	}

	sqlQuery, calculator, err := operation.buildQuery()
	if err != nil {
		if walletErr, ok := err.(WalletError); ok {
			statusCode := getHTTPStatusForError(walletErr.Code)
			sendErrorResponse(w, walletErr.Code, walletErr.Message, statusCode)
		}
		return
	}

	amount := decimal.NewFromFloat(operation.Amount)

	newBalance, err := h.processWalletOperation(ctx, operation.WalletId, amount, calculator, sqlQuery)
	if err != nil {
		if walletErr, ok := err.(WalletError); ok {
			statusCode := getHTTPStatusForError(walletErr.Code)
			sendErrorResponse(w, walletErr.Code, walletErr.Message, statusCode)
		}
		return
	}

	sendSuccessResponse(w, "balance", newBalance.String())
}

func (h *Handler) processWalletOperation(ctx context.Context, walletID string, amount decimal.Decimal, calc balanceCalculator, query string) (decimal.Decimal, error) {
	tx, err := h.db.Begin(ctx)
	if err != nil {
		return decimal.Zero, newErrDBConnectionFailed(err)
	}

	defer tx.Rollback(ctx)
	row := tx.QueryRow(ctx, "SELECT balance FROM wallets WHERE id = $1 FOR UPDATE", walletID)

	var balance decimal.Decimal
	if err = row.Scan(&balance); err != nil {
		if err == pgx.ErrNoRows {
			return decimal.Zero, newErrWalletNotFound()
		}
		return decimal.Zero, newErrDBConnectionFailed(err)
	}

	newBalance, err := calc(balance, amount)
	if err != nil {
		return decimal.Zero, err
	}

	if newBalance.IsNegative() {
		return decimal.Zero, newErrInsufficientFunds()
	}

	if _, err = tx.Exec(ctx, query, newBalance, walletID); err != nil {
		return decimal.Zero, newErrDBConnectionFailed(err)
	}

	if err = tx.Commit(ctx); err != nil {
		return decimal.Zero, newErrDBConnectionFailed(err)
	}

	return newBalance, nil
}

func (h *Handler) GetWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletIdStr := vars["WALLET_UUID"]
	walletId, err := uuid.Parse(walletIdStr)
	if err != nil {
		sendErrorResponse(w, codeInvalidWalletId, "wallet id must be a valid UUID", http.StatusBadRequest)
		return
	}

	balance, err := h.getWalletBalance(r.Context(), walletId.String())
	if err != nil {
		if walletErr, ok := err.(WalletError); ok {
			statusCode := getHTTPStatusForError(walletErr.Code)
			if walletErr.Code == codeWalletNotFound {
				statusCode = http.StatusNotFound
			}
			sendErrorResponse(w, walletErr.Code, walletErr.Message, statusCode)
			return
		}
		sendErrorResponse(w, "INTERNAL_ERROR", "An internal error occurred", http.StatusInternalServerError)
		return
	}
	sendSuccessResponse(w, "balance", balance.String())
}

func (h *Handler) getWalletBalance(ctx context.Context, walletID string) (decimal.Decimal, error) {
	var balance decimal.Decimal

	err := h.db.QueryRow(ctx, "SELECT balance FROM wallets WHERE id = $1", walletID).Scan(&balance)
	if err != nil {
		if err == pgx.ErrNoRows {
			return decimal.Zero, newErrWalletNotFound()
		}
		return decimal.Zero, newErrDBConnectionFailed(err)
	}
	return balance, nil
}
