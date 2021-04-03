package response

import (
	"encoding/json"
	"net/http"
)

type res struct {
	RSuccess    string      `json:"success"`
	RMessage    string      `json:"message"`
	RStatusCode int         `json:"statusCode"`
	RData       interface{} `json:"data"`
}

func JSON(w http.ResponseWriter, success string, message string, statusCode int, data interface{}) error {
	w.Header().Set("content-type", "aplication/json")
	w.WriteHeader(statusCode)
	response := res{RSuccess: success, RMessage: message, RStatusCode: statusCode, RData: data}
	err := json.NewEncoder(w).Encode(&response)
	if err != nil {
		return err
	}
	return nil
}

func ERROR(w http.ResponseWriter, success string, message string, statusCode int, err error) {
	if err != nil {
		JSON(w, success, message, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, success, message, http.StatusBadRequest, nil)
}
