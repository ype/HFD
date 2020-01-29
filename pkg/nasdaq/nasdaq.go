package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type NasdaqAPI struct {
	Data struct {
		Symbol         string `json:"symbol"`
		CompanyName    string `json:"companyName"`
		StockType      string `json:"stockType"`
		Exchange       string `json:"exchange"`
		IsNasdaqListed bool   `json:"isNasdaqListed"`
		IsNasdaq100    bool   `json:"isNasdaq100"`
		IsHeld         bool   `json:"isHeld"`
		PrimaryData    struct {
			LastSalePrice      string `json:"lastSalePrice"`
			NetChange          string `json:"netChange"`
			PercentageChange   string `json:"percentageChange"`
			DeltaIndicator     string `json:"deltaIndicator"`
			LastTradeTimestamp string `json:"lastTradeTimestamp"`
			IsRealTime         bool   `json:"isRealTime"`
		} `json:"primaryData"`
		SecondaryData struct {
			LastSalePrice      string `json:"lastSalePrice"`
			NetChange          string `json:"netChange"`
			PercentageChange   string `json:"percentageChange"`
			DeltaIndicator     string `json:"deltaIndicator"`
			LastTradeTimestamp string `json:"lastTradeTimestamp"`
			IsRealTime         bool   `json:"isRealTime"`
		} `json:"secondaryData"`
		KeyStats struct {
			Volume struct {
				Label string `json:"label"`
				Value string `json:"value"`
			} `json:"Volume"`
			PreviousClose struct {
				Label string `json:"label"`
				Value string `json:"value"`
			} `json:"PreviousClose"`
			OpenPrice struct {
				Label string `json:"label"`
				Value string `json:"value"`
			} `json:"OpenPrice"`
			MarketCap struct {
				Label string `json:"label"`
				Value string `json:"value"`
			} `json:"MarketCap"`
		} `json:"keyStats"`
		MarketStatus     string `json:"marketStatus"`
		AssetClass       string `json:"assetClass"`
		ComplianceStatus string `json:"complianceStatus"`
	} `json:"data"`
	Message interface{} `json:"message"`
	Status  struct {
		RCode            int         `json:"rCode"`
		BCodeMessage     interface{} `json:"bCodeMessage"`
		DeveloperMessage interface{} `json:"developerMessage"`
	} `json:"status"`
}

func main() {
	symbol := "AAPL"

	// QueryEscape escapes the symbol string so
	// it can be safely used inside a URL query
	safeSymbol := url.QueryEscape(symbol)

	url := fmt.Sprintf("https://api.nasdaq.com/api/quote/%s/info?assetclass=stocks", safeSymbol)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()

	var record NasdaqAPI

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	fmt.Println("Symbol: ", record.Data.Symbol)
	fmt.Println("Open: ", record.Data.KeyStats.OpenPrice)
	fmt.Println("Close: ", record.Data.KeyStats.PreviousClose)
	fmt.Println("Volumes: ", record.Data.KeyStats.Volume)
	fmt.Println("Last: ", record.Data.PrimaryData.LastSalePrice)
}
