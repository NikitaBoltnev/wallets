package handler

import (
	"math"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
)

type balanceCalculator func(balance, amount decimal.Decimal) (decimal.Decimal, error)

type WalletOperationRequest struct {
	WalletId      string  `json:"valletId"`
	OperationType string  `json:"operationType"`
	Amount        float64 `json:"amount"`
}

func (w *WalletOperationRequest) validation() error {
	if w.Amount <= 0 {
		return newErrAmountMustBePositive()
	}

	scaledAmount := w.Amount * 100
	roundedScaledAmount := math.Round(scaledAmount)
	if math.Abs(scaledAmount-roundedScaledAmount) > 1e-9 {
		return newErrAmountPrecisionInvalid()
	}

	if w.OperationType != "DEPOSIT" && w.OperationType != "WITHDRAW" {
		return newErrInvalidOperationType()
	}
	if w.WalletId == "" {
		return newErrWalletIdEmpty()
	}
	if _, err := uuid.Parse(w.WalletId); err != nil {
		return newErrInvalidWalletId()
	}

	return nil
}

func (w *WalletOperationRequest) buildQuery() (string, balanceCalculator, error) {
	sql := "UPDATE wallets SET balance = $1 WHERE id = $2;"
	switch w.OperationType {
	case "DEPOSIT":
		calculator := func(balance, amount decimal.Decimal) (decimal.Decimal, error) {
			newBalance := balance.Add(amount)
			return newBalance, nil
		}
		return sql, calculator, nil
	case "WITHDRAW":
		calculator := func(balance, amount decimal.Decimal) (decimal.Decimal, error) {
			if balance.LessThan(amount) {
				return decimal.Zero, newErrInsufficientFunds()
			}
			newBalance := balance.Sub(amount)
			return newBalance, nil
		}
		return sql, calculator, nil
	default:
		return "", nil, newErrInvalidOperationType()
	}
}

type Handler struct {
	db *pgxpool.Pool
}

func NewHandler(dbPool *pgxpool.Pool) *Handler {
	return &Handler{
		db: dbPool,
	}
}
