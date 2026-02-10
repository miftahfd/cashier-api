package handlers

import (
	"cashier-api/models"
	"cashier-api/response"
	"cashier-api/services"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := r.URL.Query().Get("name")

	products, err := h.service.GetAll(name)
	if err != nil {
		response.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := response.ResponseWithData{
		Status:  true,
		Message: "Get All Product",
		Data:    products,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		response.ErrorResponse(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = h.service.Create(&product)
	if err != nil {
		response.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := response.ResponseWithData{
		Status:  true,
		Message: "Create Product",
		Data:    product,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *ProductHandler) HandleProductByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetByID(w, r)
	case http.MethodPut:
		h.Update(w, r)
	case http.MethodDelete:
		h.Delete(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.service.GetByID(id)
	if err != nil {
		response.ErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	response := response.ResponseWithData{
		Status:  true,
		Message: "Get Product",
		Data:    product,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product models.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		response.ErrorResponse(w, "Invalid request", http.StatusBadRequest)
		return
	}

	product.ID = id
	err = h.service.Update(&product)
	if err != nil {
		response.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := response.ResponseWithData{
		Status:  true,
		Message: "Update Product",
		Data:    product,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		response.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := response.Response{
		Status:  true,
		Message: "Success delete product",
	}
	json.NewEncoder(w).Encode(response)
}
