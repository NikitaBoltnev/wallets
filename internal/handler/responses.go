package handler

import (
	"encoding/json"
	"net/http"
)

func sendSuccessResponse(w http.ResponseWriter, dataType string, dataValue string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	responseBody := struct {
		Success bool              `json:"success"`
		Data    map[string]string `json:"data"`
	}{
		Success: true,
		Data:    map[string]string{dataType: dataValue},
	}
	json.NewEncoder(w).Encode(responseBody)
}

func sendErrorResponse(w http.ResponseWriter, code, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	responseErrorBody := struct {
		Success bool         `json:"success"`
		Error   *WalletError `json:"error"`
	}{
		Success: false,
		Error: &WalletError{
			Code:    code,
			Message: message,
		},
	}
	json.NewEncoder(w).Encode(responseErrorBody)
}
