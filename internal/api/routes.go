package api

import (
    "github.com/gorilla/mux"
)

func SetupRoutes(h *Handler) *mux.Router {
    r := mux.NewRouter()

    r.HandleFunc("/receipts/process", h.ProcessReceipt).Methods("POST")
    r.HandleFunc("/receipts/{id}/points", h.GetPoints).Methods("GET")

    return r
}
