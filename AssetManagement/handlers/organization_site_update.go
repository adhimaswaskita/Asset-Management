package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	"github.com/gorilla/mux"
)

//UpdateOrganizationSite is handler used to update OrganizationSite by ID
func (h Handler) UpdateOrganizationSite(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	organizationSiteParam := &nmodel.OrganizationSite{}

	params := mux.Vars(r)
	ID := params["id"]

	intID, _ := strconv.Atoi(ID)
	uintID := uint(intID)

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(organizationSiteParam)

	result, err := h.Service.UpdateOrganizationSite(uintID, organizationSiteParam)
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, result, nil, w)
}
