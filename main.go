package main

import (
	"fmt"
)

func main() {

	apiKey, err := GetKey()
	if err != nil {
		fmt.Println("error during key retrieval", err)
	}

	url, err := BuildURLWithAPIKey("AAPL", "2023-08-10", "2023-08-10", apiKey)
	fmt.Println(url)
	if err != nil {
		fmt.Println("can't format url", err)
	}

	responseStream, err := MakeAPIRequest(url, apiKey)
	if err != nil {
		fmt.Println("error querying api", err)
	}

	response, err := ParseAPIResponse(responseStream)

	results := response.Results

	for i := range len(results) {
		fmt.Println("volume: ", results[i].Volume)
	}
	for _, r := range results {
		fmt.Println("high", r.High, "low", r.Low, "date", r.Timestamp)
	}
}
