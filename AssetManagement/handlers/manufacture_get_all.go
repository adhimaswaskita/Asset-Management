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
		rf.Response(nrf.ERROR, nil, w)
		return
	}

	rf.Response(nrf.SUCCESS, manufactures, w)
	return
}
