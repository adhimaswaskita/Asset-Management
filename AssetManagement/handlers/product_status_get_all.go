package handlers

import (
	"net/http"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
)

//GetAllProductStatus is used to get all product Status data
func (h *Handler) GetAllProductStatus(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}

	result, err := h.Service.GetAllProductStatus()
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, result, nil, w)
	return
}
