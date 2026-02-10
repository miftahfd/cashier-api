package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type ResponseWithData struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ErrorResponse(w http.ResponseWriter, message string, httpStatus int) {
	response := Response{
		Status:  false,
		Message: message,
	}
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(response)
}
