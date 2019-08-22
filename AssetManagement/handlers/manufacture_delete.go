package handlers

import (
	"net/http"
	"strconv"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	"github.com/gorilla/mux"
)

//DeleteManufacture is delete manufacture handler
func (h *Handler) DeleteManufacture(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	manufactureParam := &nmodel.Manufacture{}

	params := mux.Vars(r)
	id := params["id"]

	intID, _ := strconv.Atoi(id)
	ID := uint(intID)
	err := h.Service.DeleteManufacture(ID)
	if err != nil {
		rf.Response(nrf.ERROR, manufactureParam, w)
		return
	}

	rf.Response(nrf.SUCCESS, manufactureParam, w)
}
