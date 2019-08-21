package handlers

import (
	"encoding/json"
	"net/http"

	nrf "github.com/adhimaswaskita/netmonk-asset-management/lib/responseformat"
	nmodel "github.com/adhimaswaskita/netmonk-asset-management/models"
)

//CreateProductType is used to create new product type
func (h *Handler) CreateProductType(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	productTypeParam := &nmodel.ProductType{}

	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(productTypeParam)

	productType, err := h.Service.CreateProductType(productTypeParam)
	if err != nil {
		rf.Response(nrf.ERROR, productType, w)
		return
	}

	rf.Response(nrf.SUCCESS, productType, w)
}
