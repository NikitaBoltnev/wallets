package main

import (
	"log"
	"net/http"

	"github.com/NikitaBoltnev/wallets/config"
	"github.com/NikitaBoltnev/wallets/internal/handler"
	"github.com/gorilla/mux"
)

func main() {
	cfg, pool, err := config.LoadConfigAndPool()
	if err != nil {
		log.Fatal("Failed to initialize config or database pool: ", err)
	}

	h := handler.NewHandler(pool)

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/wallet", h.PostWalletTransaction).Methods("POST")
	r.HandleFunc("/api/v1/wallets/{WALLET_UUID}", h.GetWallet).Methods("GET")

	serverPort := cfg.Port
	address := ":" + serverPort
	log.Fatal(http.ListenAndServe(address, r))
}
