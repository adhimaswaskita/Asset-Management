package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"

	nhandler "github.com/adhimaswaskita/AssetManagement/handlers"
)

func (m *MockedService) DeleteManufacture(ID uint) error {
	if m.ErrMap["ErrorDeleteManufacture"] {
		return fmt.Errorf("%v", m.ErrStatement)
	}

	return nil
}

func TestDeleteManufacture(t *testing.T) {
	service := &MockedService{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall error",
	}
	handler := nhandler.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/manufacture", handler.DeleteManufacture).Methods("DELETE")
	server := httptest.NewServer(r)
	defer server.Close()

	t.Run("Delete Manufacture OK", func(t *testing.T) {
		ID := 1
		stringID := strconv.Itoa(ID)
		req, _ := http.NewRequest("DELETE", server.URL+"/manufacture?ID="+stringID, nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("%v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Status code want %v but have %v", http.StatusOK, resp.StatusCode)
		}

		values := req.URL.Query()["ID"]
		if len(values[0]) < 1 {
			t.Errorf("Parameter doesn't contain any ID")
		}

		if values[0] != "1" {
			t.Errorf("Parameter want 1 but have %v", values[0])
		}
	})

	t.Run("Create Manufacture NOK", func(t *testing.T) {
		service.ErrMap["ErrorDeleteManufacture"] = true
		req, _ := http.NewRequest("DELETE", server.URL+"/manufacture", nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("%v", resp.StatusCode)
		}

		if resp.StatusCode == http.StatusOK {
			t.Errorf("This should be 400 but have: %v", resp.StatusCode)
		}
	})
}