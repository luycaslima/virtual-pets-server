package responses

import (
	"encoding/json"
	"net/http"
)

type ResponseModel struct {
	StatusCode int         `json:"statusCode"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func EncodeResponse(rw http.ResponseWriter, statusCode int, message string, success bool, data interface{}) {
	rw.WriteHeader(statusCode)
	response := ResponseModel{
		StatusCode: statusCode,
		Message:    message,
		Success:    success,
		Data:       data,
	}
	json.NewEncoder(rw).Encode(response)
}
