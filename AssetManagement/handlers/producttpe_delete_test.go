package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"

	nhandler "github.com/adhimaswaskita/netmonk-asset-management/handlers"
)

func (m *MockedService) DeleteProductType(ID int) error {
	if m.ErrMap["ErrorCreateAllProductType"] {
		return fmt.Errorf("%v", m.ErrStatement)
	}

	return nil
}

func TestDeleteProductType(t *testing.T) {
	service := &MockedService{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall error",
	}
	handler := nhandler.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/producttype", handler.DeleteProductType).Methods("DELETE")
	server := httptest.NewServer(r)
	defer server.Close()

	t.Run("Delete product type OK", func(t *testing.T) {
		ID := 1
		stringID := strconv.Itoa(ID)
		req, _ := http.NewRequest("DELETE", server.URL+"/producttype?ID="+stringID, nil)
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

	t.Run("Create Product Type NOK", func(t *testing.T) {
		service.ErrMap["ErrorCreateAllProductType"] = true
		req, _ := http.NewRequest("DELETE", server.URL+"/producttype", nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("%v", resp.StatusCode)
		}

		if resp.StatusCode == http.StatusOK {
			t.Errorf("This should be 400 but have: %v", resp.StatusCode)
		}
	})
}
