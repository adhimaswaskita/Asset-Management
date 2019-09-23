package handlers

import (
	"encoding/json"
	"net/http"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//CreateProductType is used to create new product type
func (h *Handler) CreateProductType(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	productTypeParam := &nmodel.ProductType{}

	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(productTypeParam)

	result, err := h.Service.CreateProductType(productTypeParam)
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, result, nil, w)
}
