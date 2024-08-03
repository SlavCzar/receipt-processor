package repository

import (
    "testing"
    "receipt-processor/internal/models"
)

func TestSaveAndGetReceipt(t *testing.T) {
    repo := NewReceiptRepository()
    receipt := models.Receipt{Retailer: "Test Store"}

    id, err := repo.Save(receipt)
    if err != nil {
        t.Fatalf("Failed to save receipt: %v", err)
    }

    savedReceipt, err := repo.GetByID(id)
    if err != nil {
        t.Fatalf("Failed to get receipt: %v", err)
    }

    if savedReceipt.Retailer != receipt.Retailer {
        t.Errorf("Retrieved receipt does not match saved receipt")
    }
}

func TestGetNonExistentReceipt(t *testing.T) {
    repo := NewReceiptRepository()
    _, err := repo.GetByID("non-existent-id")
    if err == nil {
        t.Error("Expected error when getting non-existent receipt, got nil")
    }
}
