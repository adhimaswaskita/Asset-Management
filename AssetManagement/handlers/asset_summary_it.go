package handlers

import (
	"net/http"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
)

//ITAssetSummary is used to get all Asset  data
func (h *Handler) ITAssetSummary(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}

	assets, err := h.Service.ITAssetSummary()
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, assets, nil, w)
	return
}