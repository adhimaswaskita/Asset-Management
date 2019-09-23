package handlers

import (
	"encoding/json"
	"net/http"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//CreateProductStatus is used to create new product Status
func (h *Handler) CreateProductStatus(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	productStatusParam := &nmodel.ProductStatus{}

	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(productStatusParam)

	result, err := h.Service.CreateProductStatus(productStatusParam)
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, result, nil, w)
}
