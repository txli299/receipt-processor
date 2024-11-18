package models

type Item struct {
	ShortDescription string `json:"shortDescription" validate:"required,shortDescriptionPattern"`
	Price            string `json:"price" validate:"required,dollarAmount"`
}
