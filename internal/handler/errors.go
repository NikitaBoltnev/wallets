package handler

import (
	"fmt"
	"net/http"
)

const (
	codeInsufficientFunds      = "INSUFFICIENT_FUNDS"
	codeWalletNotFound         = "WALLET_NOT_FOUND"
	codeInvalidOperationType   = "INVALID_OPERATION_TYPE"
	codeInvalidWalletId        = "INVALID_WALLET_ID"
	codeWalletIdEmpty          = "WALLET_ID_EMPTY"
	codeAmountMustBePositive   = "AMOUNT_MUST_BE_POSITIVE"
	codeAmountPrecisionInvalid = "AMOUNT_PRECISION_INVALID"
	codeDBConnectionFailed     = "DB_CONNECTION_FAILED"
	codeInvaliJSON             = "INVALID_JSON"
)

type WalletError struct {
	Code    string
	Message string
}

func (e WalletError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

var (
	newErrInsufficientFunds = func() WalletError {
		return WalletError{Code: codeInsufficientFunds, Message: "the wallet has insufficient funds for this operation"}
	}
	newErrWalletNotFound = func() WalletError {
		return WalletError{Code: codeWalletNotFound, Message: "the specified wallet does not exist"}
	}
	newErrInvalidOperationType = func() WalletError {
		return WalletError{Code: codeInvalidOperationType, Message: "operation type must be DEPOSIT or WITHDRAW"}
	}
	newErrInvalidWalletId = func() WalletError {
		return WalletError{Code: codeInvalidWalletId, Message: "wallet id must be a valid UUID"}
	}
	newErrWalletIdEmpty = func() WalletError {
		return WalletError{Code: codeWalletIdEmpty, Message: "wallet id cannot be empty"}
	}
	newErrAmountMustBePositive = func() WalletError {
		return WalletError{Code: codeAmountMustBePositive, Message: "amount must be greater than 0"}
	}
	newErrAmountPrecisionInvalid = func() WalletError {
		return WalletError{Code: codeAmountPrecisionInvalid, Message: "amount must have at most 2 decimal places"}
	}
	newErrDBConnectionFailed = func(err error) WalletError {
		return WalletError{
			Code:    codeDBConnectionFailed,
			Message: fmt.Sprintf("failed to connect to the database: %v", err),
		}
	}
	NewErrInvaliJSON = func() WalletError {
		return WalletError{Code: codeInvaliJSON, Message: "invalid JSON"}
	}
)

func getHTTPStatusForError(errorCode string) int {
	switch errorCode {
	case codeInsufficientFunds, codeWalletNotFound, codeInvalidOperationType,
		codeInvalidWalletId, codeWalletIdEmpty, codeAmountMustBePositive,
		codeInvaliJSON, codeAmountPrecisionInvalid:
		//400
		return http.StatusBadRequest

	case codeDBConnectionFailed:
		//500
		return http.StatusInternalServerError

	default:
		//500
		return http.StatusInternalServerError
	}
}
