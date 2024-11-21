package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Config struct {
	PolygonAPIKey string `yaml:"polygon-api-key"`
}

type APIResponse struct {
	Ticker       string   `json:"ticker"`
	QueryCount   int      `json:"queryCount"`
	ResultsCount int      `json:"resultsCount"`
	Adjusted     bool     `json:"adjusted"`
	Results      []Result `json:"results"`
	Status       string   `json:"status"`
	RequestID    string   `json:"request_id"`
}

type Result struct {
	Volume         int     `json:"v"`
	Open           float64 `json:"o"`
	Close          float64 `json:"c"`
	High           float64 `json:"h"`
	Low            float64 `json:"l"`
	Timestamp      int64   `json:"t"`
	NumberOfTrades int     `json:"n"`
}

func GetKey() (string, error) {
	file, err := os.Open("resources/config.yaml")
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("Error occurred trying to close file handle", err)
		}
	}(file)

	configFile, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Println("Error parsing config file:", err)
		return "", err
	}

	return config.PolygonAPIKey, nil
}

func BuildURLWithAPIKey(ticker, fromDate, toDate, apiKey string) (string, error) {
	baseURL := "https://api.polygon.io/v2/aggs/ticker/"
	endpoint := fmt.Sprintf("%s%s/range/1/day/%s/%s", baseURL, ticker, fromDate, toDate)

	// Parse the URL to add query parameters
	parsedURL, err := url.Parse(endpoint)
	if err != nil {
		return "", fmt.Errorf("error parsing URL: %w", err)
	}

	// Add the API key as a query parameter
	query := parsedURL.Query()
	query.Set("apiKey", apiKey)

	parsedURL.RawQuery = query.Encode()

	return parsedURL.String(), nil
}

func MakeAPIRequest(url string, apiKey string) ([]byte, error) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}

	// Set the Authorization header
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Create an HTTP client and make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("error status", resp.Status)
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("received non-OK HTTP status %s: %s", resp.Status, string(bodyBytes))
	}

	// Read the response body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return bodyBytes, nil
}

func ParseAPIResponse(data []byte) (APIResponse, error) {
	var apiResponse APIResponse
	err := json.Unmarshal(data, &apiResponse)
	if err != nil {
		return apiResponse, fmt.Errorf("error parsing JSON response: %w", err)
	}
	return apiResponse, nil
}
