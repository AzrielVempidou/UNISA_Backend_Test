package utils

import (
	"encoding/json"
	"net/http"
)

// CreateResponse mengirimkan respons JSON dengan format standar
func CreateResponse(w http.ResponseWriter, status bool, statusCode int, message string, data interface{}) {
	response := map[string]interface{}{
		"status":      status,
		"Status_Code": statusCode,
		"message":     message,
		"data":        data,
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// ErrorResponse mengirimkan respons kesalahan JSON dengan format standar
func ErrorResponse(w http.ResponseWriter, statusCode int, errorMessage string) {
	CreateResponse(w, false, statusCode, "Terjadi kesalahan: "+errorMessage, nil)
}
