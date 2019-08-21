package handlers

import (
	"net/http"

	nrf "github.com/adhimaswaskita/netmonk-asset-management/lib/responseformat"
)

//GetAllProductType is used to get all product type data
func (h *Handler) GetAllProductType(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}

	productTypes, err := h.Service.GetAllProductType()
	if err != nil {
		rf.Response(nrf.ERROR, nil, w)
		return
	}

	rf.Response(nrf.SUCCESS, productTypes, w)
	return
}
