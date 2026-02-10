package handlers

import (
	"cashier-api/models"
	"cashier-api/response"
	"cashier-api/services"
	"encoding/json"
	"net/http"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Checkout(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TransactionHandler) HandleReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetReport(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.ErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	transaction, err := h.service.Checkout(req.Items, true)
	if err != nil {
		response.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := response.ResponseWithData{
		Status:  true,
		Message: "Checkout",
		Data:    transaction,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *TransactionHandler) GetReport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	report, err := h.service.GetReport()
	if err != nil {
		response.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := response.ResponseWithData{
		Status:  true,
		Message: "Get Report",
		Data:    report,
	}
	json.NewEncoder(w).Encode(response)
}
