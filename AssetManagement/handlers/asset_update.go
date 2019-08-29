package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//UpdateAsset is used to update Asset
func (h *Handler) UpdateAsset(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	assetParam := &nmodel.Asset{}

	params := mux.Vars(r)
	id := params["id"]

	intID, _ := strconv.Atoi(id)
	ID := uint(intID)

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(assetParam)

	Asset, err := h.Service.UpdateAsset(ID, assetParam)
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, Asset, nil, w)
}
