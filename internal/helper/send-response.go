package helper

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/miriam-samuels/portfolio-builder/internal/types"
)

func SendResponse(w http.ResponseWriter, statusCode int, status bool, message string, data interface{}, err ...error) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization")
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(types.Response{
		Status:  status,
		Message: message,
		Data:    data,
	})

	if status == false && len(err) > 0 {
		log.Printf(message+":: %v", err[0].Error())
	}
}
