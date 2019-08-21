package responseformat

import (
	"encoding/json"
	"net/http"
)

const (
	//SUCCESS is default jsend succes string
	SUCCESS = "success"
	//FAILED is default jsend succes string
	FAILED = "failed"
	//ERROR is default jsend succes string
	ERROR = "error"
)

//ResponseFormat is json response format that implement jsend standard
type ResponseFormat struct {
	Status string      `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

//Response is response format function that implement jsend standard
func (rf *ResponseFormat) Response(status string, data interface{}, w http.ResponseWriter) {
	rf.Status = status
	rf.Data = data
	w.Header().Set("Content-Type", "application/json")

	resp, err := json.Marshal(rf)
	if err != nil {
		resErr := ResponseFormat{
			Status: ERROR,
		}
		jsonErr, _ := json.Marshal(resErr)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonErr)
	}

	if status == SUCCESS {
		w.WriteHeader(http.StatusOK)
	} else if status == FAILED {
		w.WriteHeader(http.StatusBadRequest)
	} else if status == ERROR {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(resp)
}
