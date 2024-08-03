package utils

import (

    "math"
    "regexp"
    "strconv"
    "strings"
    "time"

    "receipt-processor/internal/models"
)

func CalculatePoints(receipt models.Receipt) int {
    points := 0

    // Rule 1: One point for every alphanumeric character in the retailer name
    if receipt.Retailer != "" {
        alphanumeric := regexp.MustCompile(`[a-zA-Z0-9]`)
		retailerName := strings.ReplaceAll(receipt.Retailer, " ", "")
        points += len(alphanumeric.FindAllString(retailerName, -1))
    }

    // Rule 2: 50 points if the total is a round dollar amount with no cents
    if receipt.Total != "" {
		// Check if the total ends with ".00" assuming total provided will always be a float
		if strings.HasSuffix(receipt.Total, ".00") {
			points += 50
		}
	}

    // Rule 3: 25 points if the total is a multiple of 0.25
    if receipt.Total != "" {
        total, _ := strconv.ParseFloat(receipt.Total, 64)
        if math.Mod(total*100, 25) == 0 {
            points += 25
        }
    }

    // Rule 4: 5 points for every two valid items on the receipt
    validItems := 0
    for _, item := range receipt.Items {
        if strings.TrimSpace(item.ShortDescription) != "" && item.Price != "" {
            validItems++
        }
    }
    points += (validItems / 2) * 5

    // Rule 5: Points based on valid item descriptions
    for _, item := range receipt.Items {
        trimmedDescription := strings.TrimSpace(item.ShortDescription)
        if trimmedDescription != "" && item.Price != "" {
            trimmedLength := len(trimmedDescription)
            price, _ := strconv.ParseFloat(item.Price, 64)
            if trimmedLength%3 == 0 {
                points += int(math.Ceil(price * 0.2))
            }
        }
    }

    // Rule 6: 6 points if the day in the purchase date is odd
    if receipt.PurchaseDate != "" {
        purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
        if purchaseDate.Day()%2 != 0 {
            points += 6
        }
    }

    // Rule 7: 10 points if the time of purchase is after 2:00pm and before 4:00pm including both
    if receipt.PurchaseTime != "" {
        purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
        if purchaseTime.Hour() >= 14 && purchaseTime.Hour() <= 16 {
            points += 10
        }
    }

    return points
}