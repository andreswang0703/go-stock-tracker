package internal

import (
	"context"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
	"os"
)

func GetPreviousCloseForTickers(tickers []string) ([]*models.GetPreviousCloseAggResponse, error) {
	var res []*models.GetPreviousCloseAggResponse
	for _, ticker := range tickers {
		response, err := GetPreviousClose(ticker)
		if err != nil {
			return res, err
		}
		res = append(res, response)
	}
	return res, nil
}

func GetPreviousClose(ticker string) (*models.GetPreviousCloseAggResponse, error) {
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	params := models.GetPreviousCloseAggParams{
		Ticker: ticker,
	}.WithAdjusted(true)

	agg, err := c.GetPreviousCloseAgg(context.Background(), params)
	if err != nil {
		return &models.GetPreviousCloseAggResponse{}, err
	}
	return agg, nil
}
