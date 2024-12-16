package main

import (
	"fmt"
	"net/http"

	"github.com/novychok/trustwallet"
	"github.com/novychok/trustwallet/handler"
	"github.com/novychok/trustwallet/repository"
	"github.com/novychok/trustwallet/service"
)

func main() {

	repository := repository.NewMapRepository()
	service := service.New(repository)
	handler := handler.New(service)

	// Get parsed block
	http.HandleFunc("/get-block", handler.GetBlock)

	// Get address transactions
	http.HandleFunc("/get-transactions", handler.GetTransactions)

	srv := new(trustwallet.Server)
	if err := srv.Run("8911"); err != nil {
		fmt.Printf("error to run server on port 8911: err = %s\n", err)
	}
}
