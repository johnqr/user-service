package http

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func BadRequest(w http.ResponseWriter, message string) {
	JSON(w, http.StatusBadRequest, ErrorResponse{Error: message})
}

func NotFound(w http.ResponseWriter, message string) {
	JSON(w, http.StatusNotFound, ErrorResponse{Error: message})
}

func InternalError(w http.ResponseWriter, message string) {
	JSON(w, http.StatusInternalServerError, ErrorResponse{Error: message})
}
