package handlers

import (
	"encoding/json"
	"net/http"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//CreateOrganization is handler used to create Organization
func (h *Handler) CreateOrganization(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	organizationParam := &nmodel.Organization{}

	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(organizationParam)

	result, err := h.Service.CreateOrganization(organizationParam)
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, result, nil, w)
}
