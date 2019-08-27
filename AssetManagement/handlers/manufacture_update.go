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
	id := params["id"]

	intID, _ := strconv.Atoi(id)
	ID := uint(intID)

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(manufactureParam)

	manufacture, err := h.Service.UpdateManufacture(ID, manufactureParam)
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, manufacture, nil, w)
}
