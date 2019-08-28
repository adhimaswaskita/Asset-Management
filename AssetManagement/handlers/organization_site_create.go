package handlers

import (
	"encoding/json"
	"net/http"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//CreateOrganizationSite is handler used to create OrganizationSite
func (h *Handler) CreateOrganizationSite(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	organizationSiteParam := &nmodel.OrganizationSite{}

	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(organizationSiteParam)

	organizationSite, err := h.Service.CreateOrganizationSite(organizationSiteParam)
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, organizationSite, nil, w)
}
