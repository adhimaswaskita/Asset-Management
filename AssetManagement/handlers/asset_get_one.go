package handlers

import (
	"net/http"
	"strconv"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	"github.com/gorilla/mux"
)

//GetOneAsset is used to get one Asset by it's ID
func (h *Handler) GetOneAsset(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}

	params := mux.Vars(r)
	ID := params["id"]

	intID, _ := strconv.Atoi(ID)
	uintID := uint(intID)

	result, err := h.Service.GetOneAsset(uintID)
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, result, nil, w)
	return
}
