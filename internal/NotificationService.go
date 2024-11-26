package internal

import (
	"fmt"
	"github.com/polygon-io/client-go/rest/models"
	"log"
)

func NotifyClosePrice() error {
	subscribedTickers := getSubscribedTickers()
	response, err := GetPreviousCloseForTickers(subscribedTickers)
	if err != nil {
		log.Fatal("Cannot retrieve closing prices ", err)
		return err
	}
	closingPrices := extractClosingPrice(response)
	notify(closingPrices)
	return nil
}

func getSubscribedTickers() []string {
	// move to a DB
	return []string{"AAPL", "TSLA", "GOOGL", "AMZN", "MSFT"}
}

func extractClosingPrice(responses []*models.GetPreviousCloseAggResponse) map[string]float64 {
	closingPriceMap := make(map[string]float64)
	for _, response := range responses {
		ticker := response.Ticker
		closingPrice := response.Results[0].Close
		closingPriceMap[ticker] = closingPrice
	}
	return closingPriceMap
}

func notify(closingPriceMap map[string]float64) {
	// add email service
	for ticker, price := range closingPriceMap {
		fmt.Printf("Closing price of %s is %f\n", ticker, price)
	}
}
