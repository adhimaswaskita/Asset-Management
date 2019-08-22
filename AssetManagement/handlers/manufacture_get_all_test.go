package handlers_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	nhandler "github.com/adhimaswaskita/AssetManagement/handlers"
	nreadfile "github.com/adhimaswaskita/AssetManagement/lib/readfile"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	"github.com/gorilla/mux"
)

func (s *MockedService) GetAllManufacture() ([]nmodel.Manufacture, error) {
	if s.ErrMap["ErrorGetAllManufacture"] {
		return nil, fmt.Errorf("%v", s.ErrStatement)
	}

	bManufacture, err := nreadfile.OpenFile("_fixtures/manufacture.json")
	if err != nil {
		return nil, err
	}

	manufacture := []nmodel.Manufacture{}
	json.Unmarshal(bManufacture, &manufacture)

	return manufacture, nil
}

func TestGetAllManufacture(t *testing.T) {
	service := &MockedService{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall Error",
	}
	handler := nhandler.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/manufacture", handler.GetAllManufacture).Methods("GET")
	server := httptest.NewServer(r)
	defer server.Close()

	t.Run("Get All Manufacture", func(t *testing.T) {
		req, _ := http.NewRequest("GET", server.URL+"/manufacture", nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("%v", resp.StatusCode)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("This should be 200 but have: %v", resp.StatusCode)
		}

		readBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("%v", err)
		}

		manufactures := []nmodel.Manufacture{}
		encodeResponse(readBody, &manufactures)

		if manufactures[0].Name != "Manufacture 1" {
			t.Errorf("Manufacture Name be %v but have: %v", "Manufacture 1", manufactures[0].Name)
		}
	})

	t.Run("Get All Manufacture NOK", func(t *testing.T) {
		service.ErrMap["ErrorGetAllManufacture"] = true
		req, _ := http.NewRequest("GET", server.URL+"/manufacture", nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("%v", resp.StatusCode)
		}

		if resp.StatusCode == http.StatusOK {
			t.Errorf("This should be 400 but have: %v", resp.StatusCode)
		}
	})
}
