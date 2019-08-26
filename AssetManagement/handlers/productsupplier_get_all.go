package handlers

import (
	"net/http"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
)

//GetAllProductSupplier is used to get all product Supplier data
func (h *Handler) GetAllProductSupplier(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}

	productSuppliers, err := h.Service.GetAllProductSupplier()
	if err != nil {
		rf.Response(nrf.ERROR, nil, w)
		return
	}

	rf.Response(nrf.SUCCESS, productSuppliers, w)
	return
}
