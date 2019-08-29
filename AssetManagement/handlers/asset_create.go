package handlers

import (
	"encoding/json"
	"net/http"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//CreateAsset is used to create new Asset
func (h *Handler) CreateAsset(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	assetParam := &nmodel.Asset{}

	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(assetParam)

	Asset, err := h.Service.CreateAsset(assetParam)
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, Asset, nil, w)
}
