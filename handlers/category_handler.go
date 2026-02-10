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

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) HandleCategories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categories, err := h.service.GetAll()
	if err != nil {
		response.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := response.ResponseWithData{
		Status:  true,
		Message: "Get All Category",
		Data:    categories,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		response.ErrorResponse(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = h.service.Create(&category)
	if err != nil {
		response.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := response.ResponseWithData{
		Status:  true,
		Message: "Create category",
		Data:    category,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *CategoryHandler) HandleCategoryByID(w http.ResponseWriter, r *http.Request) {
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

func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	category, err := h.service.GetByID(id)
	if err != nil {
		response.ErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	response := response.ResponseWithData{
		Status:  true,
		Message: "Get Category",
		Data:    category,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	var category models.Category
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		response.ErrorResponse(w, "Invalid request", http.StatusBadRequest)
		return
	}

	category.ID = id
	err = h.service.Update(&category)
	if err != nil {
		response.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := response.ResponseWithData{
		Status:  true,
		Message: "Update Category",
		Data:    category,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		response.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := response.Response{
		Status:  true,
		Message: "Success delete category",
	}
	json.NewEncoder(w).Encode(response)
}
