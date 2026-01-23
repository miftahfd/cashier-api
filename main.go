package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type ResponseCategory struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

var categories = []Category{
	{ID: 1, Name: "Makanan", Description: "Kategori Makanan"},
	{ID: 2, Name: "Minuman", Description: "Kategori Minuman"},
	{ID: 3, Name: "Snack", Description: "Kategori Snack"},
}

func main() {
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := Response{
			Status:  true,
			Message: "API Running",
		}
		json.NewEncoder(w).Encode(response)
	})

	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
			response := ResponseCategory{
				Status:  true,
				Message: "Get Category",
				Data:    categories,
			}
			json.NewEncoder(w).Encode(response)
		} else if r.Method == "POST" {
			var newCategory Category
			err := json.NewDecoder(r.Body).Decode(&newCategory)
			if err != nil {
				errorResponse(w, "Invalid request", http.StatusBadRequest)
				return
			}

			newCategory.ID = len(categories) + 1
			categories = append(categories, newCategory)

			w.WriteHeader(http.StatusCreated)
			response := ResponseCategory{
				Status:  true,
				Message: "Create Category",
				Data:    newCategory,
			}
			json.NewEncoder(w).Encode(response)
		}
	})

	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getCategoryByID(w, r)
		} else if r.Method == "PUT" {
			updateCategory(w, r)
		} else if r.Method == "DELETE" {
			deleteCategory(w, r)
		}
	})

	fmt.Println("Server running in localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed running server")
	}
}

func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResponse(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	for _, category := range categories {
		if category.ID == id {
			response := ResponseCategory{
				Status:  true,
				Message: "Get Category",
				Data:    category,
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	errorResponse(w, "Category not found", http.StatusNotFound)
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResponse(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	var updateCategory Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		errorResponse(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i := range categories {
		if categories[i].ID == id {
			updateCategory.ID = id
			categories[i] = updateCategory

			w.WriteHeader(http.StatusOK)
			response := ResponseCategory{
				Status:  true,
				Message: "Update Category",
				Data:    updateCategory,
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	errorResponse(w, "Category not found", http.StatusNotFound)
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResponse(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	for i, cat := range categories {
		if cat.ID == id {
			categories = append(categories[:i], categories[i+1:]...)

			response := Response{
				Status:  true,
				Message: "Success delete category",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	errorResponse(w, "Category not found", http.StatusNotFound)
}

func errorResponse(w http.ResponseWriter, message string, httpStatus int) {
	response := Response{
		Status:  false,
		Message: message,
	}
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(response)
}
