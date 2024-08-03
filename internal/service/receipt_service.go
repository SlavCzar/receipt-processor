package service

import (
    "receipt-processor/internal/models"
    "receipt-processor/internal/repository"
    "receipt-processor/internal/utils"
)

type ReceiptService interface {
    ProcessReceipt(receipt models.Receipt) (string, error)
    GetPoints(id string) (int, error)
}

type receiptService struct {
    repo repository.ReceiptRepository
}

func NewReceiptService(repo repository.ReceiptRepository) ReceiptService {
    return &receiptService{repo: repo}
}

func (s *receiptService) ProcessReceipt(receipt models.Receipt) (string, error) {
    return s.repo.Save(receipt)
}

func (s *receiptService) GetPoints(id string) (int, error) {
    receipt, err := s.repo.GetByID(id)
    if err != nil {
        return 0, err
    }
    return utils.CalculatePoints(receipt), nil
}
