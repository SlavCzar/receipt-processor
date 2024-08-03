package main

import (
    "log"
    "net/http"

    "receipt-processor/internal/api"
    "receipt-processor/internal/config"
    "receipt-processor/internal/repository"
    "receipt-processor/internal/service"
)

func main() {
    cfg := config.Load()

    repo := repository.NewReceiptRepository()
    svc := service.NewReceiptService(repo)
    handler := api.NewHandler(svc)

    router := api.SetupRoutes(handler)

    log.Printf("Server starting on port %s...\n", cfg.Port)
    log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
