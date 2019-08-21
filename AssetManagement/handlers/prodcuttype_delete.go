package handlers

import (
	"net/http"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//DeleteProductType is used to delete product type by it's ID
func (h *Handler) DeleteProductType(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	productTypeParam := &nmodel.ProductType{}

	err := h.Service.DeleteProductType(1)
	if err != nil {
		rf.Response(nrf.ERROR, productTypeParam, w)
		return
	}

	rf.Response(nrf.SUCCESS, productTypeParam, w)
}
