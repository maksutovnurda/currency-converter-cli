package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const apiURL = "https://api.exchangerate-api.com/v4/latest/"

type CurrencyResponse struct {
	Rates map[string]float64 `json:"rates"`
}

func main() {
	var amount string
	var fromCurrency string
	var toCurrency string

	fmt.Print("Enter amount: ")
	fmt.Scanln(&amount)

	fmt.Print("Enter source currency (e.g., USD): ")
	fmt.Scanln(&fromCurrency)

	fmt.Print("Enter target currency (e.g., EUR): ")
	fmt.Scanln(&toCurrency)

	amountFloat, err := stringToFloat(amount)
	if err != nil {
		log.Fatal("Invalid amount. Please provide a valid number.")
	}

	rates, err := getExchangeRates()
	if err != nil {
		log.Fatal(err)
	}

	fromRate, toRate := rates[fromCurrency], rates[toCurrency]
	if fromRate == 0 || toRate == 0 {
		log.Fatal("Invalid currency code. Please check the currency codes.")
	}

	convertedAmount := amountFloat * (toRate / fromRate)
	fmt.Printf("%.2f %s is equivalent to %.2f %s\n", amountFloat, fromCurrency, convertedAmount, toCurrency)
}

func getExchangeRates() (map[string]float64, error) {
	resp, err := http.Get(apiURL + "USD")
	if err != nil {
		return nil, fmt.Errorf("network error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	var currencyResp CurrencyResponse
	err = json.Unmarshal(body, &currencyResp)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return currencyResp.Rates, nil
}

func stringToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
