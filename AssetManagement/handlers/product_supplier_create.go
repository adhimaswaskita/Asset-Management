package handlers

import (
	"encoding/json"
	"net/http"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//CreateProductSupplier is used to create new product supplier
func (h *Handler) CreateProductSupplier(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	productSupplierParam := &nmodel.ProductSupplier{}

	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(productSupplierParam)

	result, err := h.Service.CreateProductSupplier(productSupplierParam)
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, result, nil, w)
}
