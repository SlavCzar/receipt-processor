package service

import (
    "testing"
    "receipt-processor/internal/models"
    "receipt-processor/internal/repository"
)

func TestProcessReceipt(t *testing.T) {
    repo := repository.NewReceiptRepository()
    svc := NewReceiptService(repo)

    receipt := models.Receipt{Retailer: "Test Store"}

    id, err := svc.ProcessReceipt(receipt)
    if err != nil {
        t.Fatalf("Failed to process receipt: %v", err)
    }

    if id == "" {
        t.Error("Expected non-empty ID")
    }
}

func TestGetPoints(t *testing.T) {
    repo := repository.NewReceiptRepository()
    svc := NewReceiptService(repo)

    receipt := models.Receipt{
        Retailer: "Test Store",
        Total:    "100.00",
    }

    id, _ := svc.ProcessReceipt(receipt)

    points, err := svc.GetPoints(id)
    if err != nil {
        t.Fatalf("Failed to get points: %v", err)
    }

    expectedPoints := 84 // 50 for round dollar + 9 for retailer name + 25 for multiple of 0.25
    if points != expectedPoints {
        t.Errorf("Expected %d points, got %d", expectedPoints, points)
    }
}
