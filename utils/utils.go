package utils

import (
	"encoding/json"
)

func ErrorResponse(err error) string {
	response := Response{Succeed: false, Err: err.Error()}
	resp, _ := json.Marshal(response)
	return string(resp)
}