package repository

import (
    "errors"
    "receipt-processor/internal/models"
    "sync"

    "github.com/google/uuid"
)

type ReceiptRepository interface {
    Save(receipt models.Receipt) (string, error)
    GetByID(id string) (models.Receipt, error)
}

type inMemoryReceiptRepository struct {
    receipts map[string]models.Receipt
    mutex    sync.RWMutex
}

func NewReceiptRepository() ReceiptRepository {
    return &inMemoryReceiptRepository{
        receipts: make(map[string]models.Receipt),
    }
}

func (r *inMemoryReceiptRepository) Save(receipt models.Receipt) (string, error) {
    r.mutex.Lock()
    defer r.mutex.Unlock()

    id := uuid.New().String()
    receipt.ID = id
    r.receipts[id] = receipt
    return id, nil
}

func (r *inMemoryReceiptRepository) GetByID(id string) (models.Receipt, error) {
    r.mutex.RLock()
    defer r.mutex.RUnlock()

    receipt, ok := r.receipts[id]
    if !ok {
        return models.Receipt{}, errors.New("receipt not found")
    }
    return receipt, nil
}
