package handlers

import (
	"net/http"
	"strconv"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	"github.com/gorilla/mux"
)

//DeleteOrganizationRegion is delete OrganizationRegion handler
func (h *Handler) DeleteOrganizationRegion(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}

	params := mux.Vars(r)
	ID := params["id"]

	intID, _ := strconv.Atoi(ID)
	uintID := uint(intID)
	err := h.Service.DeleteOrganizationRegion(uintID)

	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, uintID, nil, w)
}
