package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	"github.com/gorilla/mux"
)

//UpdateManufacture is handler used to update manufacture by ID
func (h Handler) UpdateManufacture(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	manufactureParam := &nmodel.Manufacture{}

	params := mux.Vars(r)
	ID := params["id"]

	intID, _ := strconv.Atoi(ID)
	uintID := uint(intID)

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(manufactureParam)

	result, err := h.Service.UpdateManufacture(uintID, manufactureParam)
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, result, nil, w)
}
