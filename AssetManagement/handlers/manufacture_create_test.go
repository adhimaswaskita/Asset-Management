package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	nhandler "github.com/adhimaswaskita/AssetManagement/handlers"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func (m *MockedService) CreateManufacture(pt *nmodel.Manufacture) (*nmodel.Manufacture, error) {
	if m.ErrMap["ErrorCreateManufacture"] {
		return nil, fmt.Errorf("%v", m.ErrStatement)
	}
	manufacture := &nmodel.Manufacture{
		Model: &gorm.Model{
			ID: 1,
		},
		Name:    "Manufacture 1",
		Comment: "comment",
	}
	return manufacture, nil
}

func TestCreateManufacture(t *testing.T) {
	service := &MockedService{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall Error",
	}

	handler := nhandler.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/manufacture", handler.CreateManufacture).Methods("POST")
	server := httptest.NewServer(r)
	defer server.Close()

	t.Run("Create Manufacture OK", func(t *testing.T) {
		manufacture := &nmodel.Manufacture{
			Name:    "Manufacture 1",
			Comment: "Comment",
		}
		jsonManufacture, _ := json.Marshal(manufacture)
		req, _ := http.NewRequest("POST", server.URL+"/manufacture", bytes.NewBuffer(jsonManufacture))

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("%v", err)
		}

		readBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("%v", err)
		}

		manufactures := nmodel.Manufacture{}
		encodeResponse(readBody, &manufactures)
		if manufactures.ID != 1 {
			t.Errorf("Manufactures ID want 1 but have %v", manufactures.ID)
		}
	})

	t.Run("Create Manufacture NOK", func(t *testing.T) {
		service.ErrMap["ErrorCreateManufacture"] = true
		manufacture := &nmodel.Manufacture{
			Name:    "Manufacture 1",
			Comment: "comment",
		}
		jsonManufacture, _ := json.Marshal(manufacture)
		req, _ := http.NewRequest("POST", server.URL+"/manufacture", bytes.NewBuffer(jsonManufacture))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("%v", resp.StatusCode)
		}

		if resp.StatusCode == http.StatusOK {
			t.Errorf("This should be 400 but have: %v", resp.StatusCode)
		}
	})
}
