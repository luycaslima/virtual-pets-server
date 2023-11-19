package responses

import (
	"encoding/json"
	"net/http"
)

// Strutcture of a Response
type Response struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type PetResponse Response
type UserResponse Response

func EncodeResponse(rw http.ResponseWriter, statusCode int, message string, data map[string]interface{}) {
	rw.WriteHeader(statusCode)
	response := Response{Status: statusCode, Message: message, Data: data}
	json.NewEncoder(rw).Encode(response)
}
