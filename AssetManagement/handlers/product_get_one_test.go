package handlers_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	nhandler "github.com/adhimaswaskita/AssetManagement/handlers"
	nreadfile "github.com/adhimaswaskita/AssetManagement/lib/readfile"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	"github.com/gorilla/mux"
)

func (s *MockedService) GetOneProduct(ID uint) (*nmodel.Product, error) {
	if s.ErrMap["ErrorGetOneProduct"] {
		return nil, fmt.Errorf("%v", s.ErrStatement)
	}

	bProducts, err := nreadfile.OpenFile("_fixtures/product.json")
	if err != nil {
		return nil, err
	}

	products := nmodel.Product{}
	json.Unmarshal(bProducts, &products)

	return &products, nil
}

func TestGetOneProduct(t *testing.T) {
	service := &MockedService{
		ErrMap:       map[string]bool{},
		ErrStatement: "Intentionall Error",
	}
	handler := nhandler.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/product/{id}", handler.GetOneProduct).Methods("GET")
	server := httptest.NewServer(r)
	defer server.Close()

	t.Run("Get One Product ", func(t *testing.T) {
		ID := 1
		stringID := strconv.Itoa(ID)

		req, _ := http.NewRequest("GET", server.URL+"/product?ID="+stringID, nil)
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

		products := []nmodel.Product{}
		encodeResponse(readBody, &products)

		values := req.URL.Query()["ID"]
		if len(values[0]) < 1 {
			t.Errorf("Parameter want 1 but have %v", values[0])
		}

		fmt.Println(products[0].Name)

		if products[0].Name != "Product 1" {
			t.Errorf("Product  Name be %v but have: %v", "Product  1", products[0].Name)
		}
	})

	t.Run("Get One Product  NOK", func(t *testing.T) {
		ID := 1
		stringID := strconv.Itoa(ID)

		service.ErrMap["ErrorGetOneProduct"] = true
		req, _ := http.NewRequest("GET", server.URL+"/product?ID="+stringID, nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("%v", resp.StatusCode)
		}

		if resp.StatusCode == http.StatusOK {
			t.Errorf("This should be 400 but have: %v", resp.StatusCode)
		}
	})
}
