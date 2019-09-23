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

	result, err := h.Service.CreateManufacture(manufactureParam)
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, result, nil, w)
}
