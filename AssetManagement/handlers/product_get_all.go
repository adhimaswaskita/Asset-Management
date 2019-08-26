package handlers

import (
	"net/http"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
)

//GetAllProduct is used to get all product  data
func (h *Handler) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}

	products, err := h.Service.GetAllProduct()
	if err != nil {
		rf.Response(nrf.ERROR, nil, w)
		return
	}

	rf.Response(nrf.SUCCESS, products, w)
	return
}
