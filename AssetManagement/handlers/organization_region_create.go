package handlers

import (
	"encoding/json"
	"net/http"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//CreateOrganizationRegion is handler used to create OrganizationRegion
func (h *Handler) CreateOrganizationRegion(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	organizationRegionParam := &nmodel.OrganizationRegion{}

	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(organizationRegionParam)

	result, err := h.Service.CreateOrganizationRegion(organizationRegionParam)
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, result, nil, w)
}
