package handlers

import (
	"net/http"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
)

//GetAllOrganizationSite is get all OrganizationSite data handler
func (h *Handler) GetAllOrganizationSite(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}

	organizationSites, err := h.Service.GetAllOrganizationSite()
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, organizationSites, nil, w)
	return
}