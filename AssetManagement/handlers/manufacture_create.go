package handlers

import (
	"encoding/json"
	"net/http"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//CreateManufacture is handler used to create manufacture
func (h *Handler) CreateManufacture(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	manufactureParam := &nmodel.Manufacture{}

	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(manufactureParam)

	manufacture, err := h.Service.CreateManufacture(manufactureParam)
	if err != nil {
		rf.Response(nrf.ERROR, manufacture, w)
		return
	}

	rf.Response(nrf.SUCCESS, manufacture, w)
}
