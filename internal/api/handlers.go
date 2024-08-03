package api

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "receipt-processor/internal/models"
    "receipt-processor/internal/service"
)

type Handler struct {
    service service.ReceiptService
}

func NewHandler(service service.ReceiptService) *Handler {
    return &Handler{service: service}
}

func (h *Handler) ProcessReceipt(w http.ResponseWriter, r *http.Request) {
    var receipt models.Receipt
    if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    id, err := h.service.ProcessReceipt(receipt)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := map[string]string{"id": id}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetPoints(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    points, err := h.service.GetPoints(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    response := map[string]int{"points": points}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
