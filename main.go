package main

import (
	"go-stock-tracker/internal"
	"log"
)

func main() {
	err := internal.NotifyClosePrice()
	if err != nil {
		log.Fatal("failed to notify closing price", err)
	}
}
