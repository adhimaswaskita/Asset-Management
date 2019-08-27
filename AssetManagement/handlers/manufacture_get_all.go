package handlers

import (
	"net/http"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
)

//GetAllManufacture is get all manufacture data handler
func (h *Handler) GetAllManufacture(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}

	manufactures, err := h.Service.GetAllManufacture()
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, manufactures, nil, w)
	return
}
