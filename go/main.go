package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RegisterRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type RegisterResponse struct {
	ApiKey string `json:"api_key"`
}

type SummaryRequest struct {
	Principal float64 `json:"principal"`
	Rate      float64 `json:"rate"`
	Months    int     `json:"months"`
	StartDate string  `json:"start_date"`
}

type SummaryResponse struct {
	Periods        int     `json:"periods"`
	MonthlyPayment float64 `json:"monthly_payment"`
	TotalInterest  float64 `json:"total_interest"`
	TotalPaid      float64 `json:"total_paid"`
	PayoffDate     string  `json:"payoff_date"`
	MonthsSaved    int     `json:"months_saved"`
	InterestSaved  float64 `json:"interest_saved"`
}

func main() {
	apiBase := "http://localhost:8080/api"

	// Step 1: Register to get an API key
	registerReq := RegisterRequest{
		Name:  "Test User",
		Email: "test@example.com",
	}
	body, _ := json.Marshal(registerReq)

	resp, err := http.Post(apiBase+"/register", "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var registerResp RegisterResponse
	json.NewDecoder(resp.Body).Decode(&registerResp)
	apiKey := registerResp.ApiKey
	fmt.Printf("Your API key: %s\n", apiKey)

	// Step 2: Make a calculation using the API key
	summaryReq := SummaryRequest{
		Principal: 200000,
		Rate:      4.5,
		Months:    300,
		StartDate: "2026-03-01",
	}
	body, _ = json.Marshal(summaryReq)

	req, _ := http.NewRequest("POST", apiBase+"/summary", bytes.NewBuffer(body))
	req.Header.Set("X-API-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var summaryResp SummaryResponse
	json.NewDecoder(resp.Body).Decode(&summaryResp)

	output, _ := json.MarshalIndent(summaryResp, "", "  ")
	fmt.Println(string(output))
}
