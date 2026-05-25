package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type PriceWithTaxRequest struct {
	Product Product `json:"product"`
	TaxRate float64 `json:"taxRate"`
}

type PriceWithTaxResponse struct {
	Name         string  `json:"name"`
	BasePrice    float64 `json:"basePrice"`
	TaxRate      float64 `json:"taxRate"`
	PriceWithTax float64 `json:"priceWithTax"`
}

type AverageRequest struct {
	Numbers []int `json:"numbers"`
}

type AverageResponse struct {
	Average float64 `json:"average"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type InventoryItem struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Status   string `json:"status"`
}

type InventoryResponse struct {
	Items []InventoryItem `json:"items"`
}

func (p Product) GetPriceWithTax(taxRate float64) float64 {
	return p.Price * (1 + taxRate)
}

func CalculateAverage(numbers []int) (float64, error) {
	if len(numbers) == 0 {
		return 0.0, errors.New("cannot calculate average of an empty slice")
	}

	total := 0
	for _, number := range numbers {
		total += number
	}

	return float64(total) / float64(len(numbers)), nil
}

func getInventory() map[string]int {
	return map[string]int{
		"apple":  10,
		"banana": 5,
		"orange": 0,
	}
}

func writeJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func priceWithTaxHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, ErrorResponse{Error: "method not allowed"})
		return
	}

	var request PriceWithTaxRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid json payload"})
		return
	}

	response := PriceWithTaxResponse{
		Name:         request.Product.Name,
		BasePrice:    request.Product.Price,
		TaxRate:      request.TaxRate,
		PriceWithTax: request.Product.GetPriceWithTax(request.TaxRate),
	}

	writeJSON(w, http.StatusOK, response)
}

func averageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, ErrorResponse{Error: "method not allowed"})
		return
	}

	var request AverageRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid json payload"})
		return
	}

	average, err := CalculateAverage(request.Numbers)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, AverageResponse{Average: average})
}

func inventoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, ErrorResponse{Error: "method not allowed"})
		return
	}

	inventory := getInventory()
	items := make([]InventoryItem, 0, len(inventory))

	for name, quantity := range inventory {
		status := "in_stock"
		if quantity == 0 {
			status = "out_of_stock"
		}

		items = append(items, InventoryItem{
			Name:     name,
			Quantity: quantity,
			Status:   status,
		})
	}

	writeJSON(w, http.StatusOK, InventoryResponse{Items: items})
}

func registerRoutes() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/product/price-with-tax", priceWithTaxHandler)
	http.HandleFunc("/average", averageHandler)
	http.HandleFunc("/inventory", inventoryHandler)
}

func main() {
	product := Product{
		Name:  "Laptop",
		Price: 1200.0,
	}

	fmt.Printf("%s price with tax: %.2f\n", product.Name, product.GetPriceWithTax(0.07))

	numbers := []int{10, 20, 30, 40}
	average, err := CalculateAverage(numbers)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Average: %.2f\n", average)
	}

	emptyNumbers := []int{}
	emptyAverage, err := CalculateAverage(emptyNumbers)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Average: %.2f\n", emptyAverage)
	}

	inventory := getInventory()
	for productName, quantity := range inventory {
		if quantity == 0 {
			fmt.Printf("%s is out of stock.\n", productName)
			continue
		}

		fmt.Printf("%s has %d items.\n", productName, quantity)
	}

	registerRoutes()
	fmt.Println("HTTP server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
