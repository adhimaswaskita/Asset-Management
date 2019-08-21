package handlers_test

import (
	"encoding/json"
	"fmt"
	"testing"

	nrf "github.com/adhimaswaskita/netmonk-asset-management/lib/responseformat"
)

type MockedService struct {
	ErrStatement string
	ErrMap       map[string]bool
}

func TestNewHandler(t *testing.T) {
}

func encodeResponse(body []byte, content interface{}) error {
	rf := nrf.ResponseFormat{}
	err := json.Unmarshal(body, &rf)
	if err != nil {
		return fmt.Errorf("Unable to unmarshal body: %v", err)
	}

	bdata, _ := json.Marshal(rf.Data)
	err = json.Unmarshal(bdata, &content)
	if err != nil {
		return fmt.Errorf("Unable to unmarshal data: %v", err)
	}

	return nil
}
