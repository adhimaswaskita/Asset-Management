package responseformat

import (
	"encoding/json"
	"net/http"
)

//ResponseWriter write expected response to http.ResponseWriter
func ResponseWriter(status string, rf interface{}, w http.ResponseWriter) {
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
