package handlers

import (
	"net/http"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
)

//TotalInRepairAsset is used to get all Asset  data
func (h *Handler) TotalInRepairAsset(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}

	result, err := h.Service.TotalInRepairAsset()
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, result, nil, w)
	return
}
