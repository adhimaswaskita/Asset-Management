package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	"github.com/gorilla/mux"
)

//UpdateOrganization is handler used to update Organization by ID
func (h Handler) UpdateOrganization(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	organizationParam := &nmodel.Organization{}

	params := mux.Vars(r)
	ID := params["id"]

	intID, _ := strconv.Atoi(ID)
	uintID := uint(intID)

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(organizationParam)

	result, err := h.Service.UpdateOrganization(uintID, organizationParam)
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, result, nil, w)
}
