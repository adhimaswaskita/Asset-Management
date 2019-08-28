package handlers

import (
	"net/http"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
)

//GetAllOrganizationRegion is get all OrganizationRegion data handler
func (h *Handler) GetAllOrganizationRegion(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}

	organizationRegions, err := h.Service.GetAllOrganizationRegion()
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, organizationRegions, nil, w)
	return
}
