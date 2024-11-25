package main

import (
	"go-stock-tracker/internal"
	"log"
)

func main() {

	price, err := internal.GetPreviousClose("AAPL")
	if err != nil {
		log.Fatalf("error retrieving the price")
	}
	closePrice := price.Results[0].Close
	log.Println("AAPL close price", closePrice)
}
