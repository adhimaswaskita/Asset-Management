package handlers

import (
	"net/http"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
)

//CountAssetInStore is used to get all Asset  data
func (h *Handler) CountAssetInStore(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}

	total, err := h.Service.CountAssetInStore()
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, total, nil, w)
	return
}
