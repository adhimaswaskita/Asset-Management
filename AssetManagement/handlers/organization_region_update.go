package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	"github.com/gorilla/mux"
)

//UpdateOrganizationRegion is handler used to update OrganizationRegion by ID
func (h Handler) UpdateOrganizationRegion(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	organizationRegionParam := &nmodel.OrganizationRegion{}

	params := mux.Vars(r)
	ID := params["id"]

	intID, _ := strconv.Atoi(ID)
	uintID := uint(intID)

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(organizationRegionParam)

	result, err := h.Service.UpdateOrganizationRegion(uintID, organizationRegionParam)
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, result, nil, w)
}
