package models

type Receipt struct {
	Retailer     string `json:"retailer" validate:"required,retailerPattern"`
	PurchaseDate string `json:"purchaseDate" validate:"required"`
	PurchaseTime string `json:"purchaseTime" validate:"required"`
	Items        []Item `json:"items" validate:"required,min=1,dive"`
	Total        string `json:"total" validate:"required,dollarAmount"`
}
