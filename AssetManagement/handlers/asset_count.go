package handlers

import (
	"net/http"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
)

//CountAsset is used to get all Asset  data
func (h *Handler) CountAsset(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}

	total, err := h.Service.CountAsset()
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, total, nil, w)
	return
}
