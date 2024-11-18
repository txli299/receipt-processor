package handlers

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/txli299/receipt-processor/models"
	"github.com/txli299/receipt-processor/store"
	"github.com/txli299/receipt-processor/utils"
)

type ProcessReceiptResponse struct {
	ID string `json:"id"`
}

type PointsResponse struct {
	Points int `json:"points"`
}

var validate = validator.New()

// Initialize custom validators
func init() {
	// Validation for dollar amounts (e.g., 9.99)
	validate.RegisterValidation("dollarAmount", func(fl validator.FieldLevel) bool {
		matched, _ := regexp.MatchString(`^\d+\.\d{2}$`, fl.Field().String())
		return matched
	})

	// Validation for retailer (e.g., "M&M Corner Market")
	validate.RegisterValidation("retailerPattern", func(fl validator.FieldLevel) bool {
		matched, _ := regexp.MatchString(`^[\w\s\-&]+$`, fl.Field().String())
		return matched
	})

	// Validation for shortDescription (e.g., "Mountain Dew 12PK")
	validate.RegisterValidation("shortDescriptionPattern", func(fl validator.FieldLevel) bool {
		matched, _ := regexp.MatchString(`^[\w\s\-]+$`, fl.Field().String())
		return matched
	})
}

// ProcessReceipt handles the POST request to process a receipt.
func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the receipt struct.
	err = validate.Struct(receipt)
	if err != nil {
		http.Error(w, "Invalid receipt: "+err.Error(), http.StatusBadRequest)
		return
	}

	points := utils.CalculatePoints(receipt)
	id := uuid.New().String()
	store.Store.SaveReceipt(id, points)

	response := ProcessReceiptResponse{ID: id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetPoints handles the GET request to retrieve points for a receipt.
func GetPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	points, err := store.Store.GetPoints(id)
	if err != nil {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	response := PointsResponse{Points: points}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
