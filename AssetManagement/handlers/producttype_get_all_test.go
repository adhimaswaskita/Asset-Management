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

func (s *MockedService) GetAllProductType() ([]nmodel.ProductType, error) {
	if s.ErrMap["ErrorGetAllProductType"] {
		return nil, fmt.Errorf("%v", s.ErrStatement)
	}

	bProductTypes, err := nreadfile.OpenFile("_fixtures/producttype.json")
	if err != nil {
		return nil, err
	}

	productTypes := []nmodel.ProductType{}
	json.Unmarshal(bProductTypes, &productTypes)

	return productTypes, nil
}

func TestGetAllProductType(t *testing.T) {
	service := &MockedService{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall Error",
	}
	handler := nhandler.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/producttype", handler.GetAllProductType).Methods("GET")
	server := httptest.NewServer(r)
	defer server.Close()

	t.Run("Get All Product Type", func(t *testing.T) {
		req, _ := http.NewRequest("GET", server.URL+"/producttype", nil)
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

		productTypes := []nmodel.ProductType{}
		encodeResponse(readBody, &productTypes)

		if productTypes[0].Name != "Product type 1" {
			t.Errorf("Product Type Name be %v but have: %v", "Product type 1", productTypes[0].Name)
		}
	})

	t.Run("Get All Product Type NOK", func(t *testing.T) {
		service.ErrMap["ErrorGetAllProductType"] = true
		req, _ := http.NewRequest("GET", server.URL+"/producttype", nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("%v", resp.StatusCode)
		}

		if resp.StatusCode == http.StatusOK {
			t.Errorf("This should be 400 but have: %v", resp.StatusCode)
		}
	})
}
