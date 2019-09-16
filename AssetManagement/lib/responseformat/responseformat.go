package responseformat

import (
	"net/http"
	"reflect"
)

const (
	//SUCCESS is default jsend succes string
	SUCCESS = "success"
	//FAILED is default jsend succes string
	FAILED = "failed"
	//ERROR is default jsend succes string
	ERROR = "error"
)

//ResponseFormat is succes json response format that implement jsend standard
type ResponseFormat struct {
	Status  string      `json:"status,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

//DeleteResponse is success json response that return ID of deleted item
type DeleteResponse struct {
	ID interface{} `json:"ID"`
}

type DashboardResponse struct {
}

//Response is response format function that implement jsend standard
func (rf *ResponseFormat) Response(status string, data interface{}, message interface{}, w http.ResponseWriter) {
	rf.Status = status
	rf.Message = message
	w.Header().Set("Content-Type", "application/json")

	reflectValue := reflect.ValueOf(data)

	if reflectValue.Kind() == reflect.Uint {
		rf.Data = DeleteResponse{
			ID: data,
		}
		ResponseWriter(status, rf, w)
	} else {
		rf.Data = data
		ResponseWriter(status, rf, w)
	}
}
