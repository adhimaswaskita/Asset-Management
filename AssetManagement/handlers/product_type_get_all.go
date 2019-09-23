package handlers

import (
	"net/http"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
)

//GetAllProductType is used to get all product type data
func (h *Handler) GetAllProductType(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}

	result, err := h.Service.GetAllProductType()
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, result, nil, w)
	return
}
