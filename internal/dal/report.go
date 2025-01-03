package dal

import (
	"Library-Management-System/internal/models"
	"encoding/json"
	"net/http"
)

func SendError(status int, message string, err error, w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	errorResponse := models.Error{
		Status:  status,
		Message: message,
		Error:   err,
	}

	jsonData, err := json.Marshal(errorResponse)

	if err != nil {
		http.Error(w, `{"Error": "Internal Server Error", "Status": 500}`, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
