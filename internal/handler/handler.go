package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payment struct {
}

type HealthCheckResponse struct {
	Status string `json:"status"`
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthCheckResponse{Status: "OK"}
	json.NewEncoder(w).Encode(response)
}

func (p *Payment) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating a payment request")
}
