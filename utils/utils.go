package utils

import (
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/txli299/receipt-processor/models"
)

// CalculatePoints calculates the total points for a given receipt based on the specified rules.
func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// Rule 1: One point for every alphanumeric character in the retailer name.
	for _, char := range receipt.Retailer {
		if isAlphanumeric(char) {
			points++
		}
	}

	// Rule 2: 50 points if the total is a round dollar amount with no cents (e.g., "10.00").
	if isRoundDollarAmount(receipt.Total) {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25.
	if isMultipleOfQuarter(receipt.Total) {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt.
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: If the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer. Add the result to points.
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// Rule 6: 6 points if the day in the purchase date is odd.
	if isOddDay(receipt.PurchaseDate) {
		points += 6
	}

	// Rule 7: 10 points if the time of purchase is after 2:00 PM and before 4:00 PM.
	if isAfternoonPurchase(receipt.PurchaseTime) {
		points += 10
	}

	return points
}

// Helper function to check if a character is alphanumeric.
func isAlphanumeric(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

// Helper function to check if the total is a round dollar amount.
func isRoundDollarAmount(total string) bool {
	return strings.HasSuffix(total, ".00")
}

// Helper function to check if the total is a multiple of 0.25.
func isMultipleOfQuarter(total string) bool {
	price, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return false
	}
	return math.Mod(price, 0.25) == 0
}

// Helper function to check if the purchase date is an odd day.
func isOddDay(date string) bool {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return false
	}
	day := parsedDate.Day()
	return day%2 != 0
}

// Helper function to check if the purchase time is between 2:00 PM and 4:00 PM.
func isAfternoonPurchase(timeStr string) bool {
	parsedTime, err := time.Parse("15:04", timeStr)
	if err != nil {
		return false
	}
	hour := parsedTime.Hour()
	return hour >= 14 && hour < 16
}
