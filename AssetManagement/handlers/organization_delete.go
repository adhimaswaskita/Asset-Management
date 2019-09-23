package handlers

import (
	"net/http"
	"strconv"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	"github.com/gorilla/mux"
)

//DeleteOrganization is delete Organization handler
func (h *Handler) DeleteOrganization(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}

	params := mux.Vars(r)
	id := params["id"]

	intID, _ := strconv.Atoi(id)
	uintID := uint(intID)
	err := h.Service.DeleteOrganization(uintID)

	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, uintID, nil, w)
}
