package utils

import (
	"fmt"
	"receipt-processor/internal/models"
	"testing"
)

func TestCalculatePoints(t *testing.T) {
    tests := []struct {
        name     string
        receipt  models.Receipt
        expected int
    }{
		{
			name:"Test full payload 1",
			receipt: models.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []models.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
					{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
					{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
					{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
					{ShortDescription: "Klarbrunn 12-PK 12 FL OZ", Price: "12.00"},
				},
				Total: "35.35",
			},
			expected: 28,
		},
		{
			name:"Test full payload 2",
			receipt: models.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-03-20",
				PurchaseTime: "14:33",
				Items: []models.Item{
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
				},
				Total: "9.00",
			},
			expected: 109,
		},
        {
            name: "Round dollar amount",
            receipt: models.Receipt{
                Retailer: "Test",
                Total:    "100.00",
            },
            expected: 79, // 50 for round dollar + 4 for retailer name + 25 for multiple of 0.25
        },
        {
            name: "Odd day purchase",
            receipt: models.Receipt{
                Retailer:     "Shop",
                PurchaseDate: "2022-03-21",
            },
            expected: 10, // 4 for retailer name + 6 for odd day
        },
        {
            name: "Multiple of 0.25",
            receipt: models.Receipt{
                Total: "9.75",
            },
            expected: 25, // 25 for multiple of 0.25
        },
        {
            name: "Items count",
            receipt: models.Receipt{
				Retailer: "Test",
                Items: []models.Item{{}, {}, {}, {}},
            },
            expected: 4, // 4 for retailer. Check whether empty items are not counted towards points
        },
        {
            name: "Description length multiple of 3",
            receipt: models.Receipt{
                Items: []models.Item{
                    {ShortDescription: "ABC", Price: "2.00"},
                },
            },
            expected: 1, // ceil(2.00 * 0.2) = 1
        },
        {
            name: "Afternoon purchase",
            receipt: models.Receipt{
                PurchaseTime: "14:30",
            },
            expected: 10,
        },
		{
            name: "Missing items",
            receipt: models.Receipt{
                Retailer: "Test",
                Total:    "50.00",
				Items: []models.Item{
					{}, {}, {},{},
				},
            },
            expected: 79, // Adjust expected points based on your logic for missing items
        },
        {
            name: "No retailer and total",
            receipt: models.Receipt{
                // No fields set
            },
            expected: 0, //  No points for empty receipt
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
			fmt.Println("Receipt passed: ", tt.receipt)
            points := CalculatePoints(tt.receipt)
            if points != tt.expected {
                t.Errorf("CalculatePoints() = %v, want %v", points, tt.expected)
            }
        })
    }
}
