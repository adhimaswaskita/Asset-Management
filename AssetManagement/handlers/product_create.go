package handlers

import (
	"encoding/json"
	"net/http"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//CreateProduct is used to create new product
func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	productParam := &nmodel.Product{}

	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(productParam)

	result, err := h.Service.CreateProduct(productParam)
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, result, nil, w)
}
