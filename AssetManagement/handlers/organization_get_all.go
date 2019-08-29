package handlers

import (
	"net/http"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
)

//GetAllOrganization is get all Organization data handler
func (h *Handler) GetAllOrganization(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}

	organizations, err := h.Service.GetAllOrganization()
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, organizations, nil, w)
	return
}
