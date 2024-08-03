package api

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"

    "receipt-processor/internal/models"
    "receipt-processor/internal/repository"
    "receipt-processor/internal/service"

    "github.com/gorilla/mux"
)

func TestProcessReceiptAndGetPoints(t *testing.T) {
    repo := repository.NewReceiptRepository()
    svc := service.NewReceiptService(repo)
    handler := NewHandler(svc)

    router := mux.NewRouter()
    router.HandleFunc("/receipts/process", handler.ProcessReceipt).Methods("POST")
    router.HandleFunc("/receipts/{id}/points", handler.GetPoints).Methods("GET")

    // Step 1: Process Receipt
    receipt := models.Receipt{
        Retailer: "Test Store",
        Total:    "100.00",
    }
    body, _ := json.Marshal(receipt)

    req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(body))
    rr := httptest.NewRecorder()

    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Handler returned wrong status code for ProcessReceipt: got %v want %v", status, http.StatusOK)
    }

    var processResponse map[string]string
    json.Unmarshal(rr.Body.Bytes(), &processResponse)

    id, exists := processResponse["id"]
    if !exists {
        t.Fatalf("Response does not contain an ID")
    }

    // Step 2: Get Points
    req, _ = http.NewRequest("GET", fmt.Sprintf("/receipts/%s/points", id), nil)
    rr = httptest.NewRecorder()

    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Handler returned wrong status code for GetPoints: got %v want %v", status, http.StatusOK)
    }

    var pointsResponse map[string]int
    json.Unmarshal(rr.Body.Bytes(), &pointsResponse)

    points, exists := pointsResponse["points"]
    if !exists {
        t.Errorf("Response does not contain points")
    }

    expectedPoints := 84 // 50 for round dollar + 9 for retailer name + 25 for multiple of 0.25
    if points != expectedPoints {
        t.Errorf("Expected %d points, got %d", expectedPoints, points)
    }
}
