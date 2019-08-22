package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	nhandler "github.com/adhimaswaskita/AssetManagement/handlers"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

func (m *MockedService) UpdateManufacture(ID uint, Manufacture *nmodel.Manufacture) (*nmodel.Manufacture, error) {
	if m.ErrMap["ErrorUpdateManufacture"] {
		return nil, fmt.Errorf("%v", m.ErrStatement)
	}
	manufacture := &nmodel.Manufacture{
		Model: &gorm.Model{
			ID: 1,
		},
		Name:    "New Manufacture 1",
		Comment: "Comment manufacture 1",
	}
	return manufacture, nil
}

func TestUpdateManufacture(t *testing.T) {
	service := &MockedService{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall error",
	}
	handler := nhandler.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/manufacture", handler.UpdateManufacture).Methods("PUT")
	server := httptest.NewServer(r)
	defer server.Close()

	t.Run("Update Manufacture OK", func(t *testing.T) {
		manufacture := &nmodel.Manufacture{
			Name:    "New Manufacture 1",
			Comment: "Comment manufacture 1",
		}
		jsonManufacture, _ := json.Marshal(manufacture)

		ID := 1
		stringID := strconv.Itoa(ID)

		req, _ := http.NewRequest("PUT", server.URL+"/manufacture?ID="+stringID, bytes.NewBuffer(jsonManufacture))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("%v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("This should be 200 but have: %v", resp.StatusCode)
		}
		values := req.URL.Query()["ID"]
		if len(values[0]) < 1 {
			t.Errorf("Parameter want 1 but have %v", values[0])
		}

		if values[0] != "1" {
			t.Errorf("Parameter want 1 but have %v", values[0])
		}

		readBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("%v", err)
		}

		manufactureResponse := &nmodel.Manufacture{}
		encodeResponse(readBody, manufactureResponse)
		if manufactureResponse.Name != "New Manufacture 1" {
			t.Errorf("manufactureResponse.Name want 'New Manufacture 1' but have %v", manufactureResponse.Name)
		}
	})

	t.Run("Create Manufacture NOK", func(t *testing.T) {
		service.ErrMap["ErrorUpdateManufacture"] = true
		manufacture := &nmodel.Manufacture{
			Name:    "New Manufacture 1",
			Comment: "comment manufacture 1",
		}
		jsonManufacture, _ := json.Marshal(manufacture)
		req, _ := http.NewRequest("PUT", server.URL+"/manufacture", bytes.NewBuffer(jsonManufacture))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("%v", resp.StatusCode)
		}

		if resp.StatusCode == http.StatusOK {
			t.Errorf("This should be 400 but have: %v", resp.StatusCode)
		}
	})
}
