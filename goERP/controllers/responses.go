package controllers

import (
	"encoding/json"
	"goERP/types/restResponses"
	"log"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, statusCode int, error restResponses.Error) {
	if error.Status == http.StatusNotFound {
		SendResponse(w, statusCode, nil)
		return
	}

	SendResponse(w, statusCode, error)
}

func SendResponse(w http.ResponseWriter, statusCode int, any interface{}) {
	w.WriteHeader(statusCode)
	if any != nil {
		if err := json.NewEncoder(w).Encode(any); err != nil {
			log.Println("Error while writing header: ", err.Error(), statusCode)
		}
	}
}
