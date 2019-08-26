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

func (s *MockedService) GetAllProductSupplier() ([]nmodel.ProductSupplier, error) {
	if s.ErrMap["ErrorGetAllProductSupplier"] {
		return nil, fmt.Errorf("%v", s.ErrStatement)
	}

	bProductSuppliers, err := nreadfile.OpenFile("_fixtures/productsupplier.json")
	if err != nil {
		return nil, err
	}

	productSuppliers := []nmodel.ProductSupplier{}
	json.Unmarshal(bProductSuppliers, &productSuppliers)

	return productSuppliers, nil
}

func TestGetAllProductSupplier(t *testing.T) {
	service := &MockedService{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall Error",
	}
	handler := nhandler.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/productsupplier", handler.GetAllProductSupplier).Methods("GET")
	server := httptest.NewServer(r)
	defer server.Close()

	t.Run("Get All Product Supplier", func(t *testing.T) {
		req, _ := http.NewRequest("GET", server.URL+"/productsupplier", nil)
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

		productSuppliers := []nmodel.ProductSupplier{}
		encodeResponse(readBody, &productSuppliers)

		if productSuppliers[0].Name != "Product Supplier 1" {
			t.Errorf("Product Supplier Name be %v but have: %v", "Product Supplier 1", productSuppliers[0].Name)
		}
	})

	t.Run("Get All Product Supplier NOK", func(t *testing.T) {
		service.ErrMap["ErrorGetAllProductSupplier"] = true
		req, _ := http.NewRequest("GET", server.URL+"/productsupplier", nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("%v", resp.StatusCode)
		}

		if resp.StatusCode == http.StatusOK {
			t.Errorf("This should be 400 but have: %v", resp.StatusCode)
		}
	})
}
