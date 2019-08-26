package handlers

import (
	"net/http"
	"strconv"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	"github.com/gorilla/mux"
)

//DeleteProduct is used to delete product  by it's ID
func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	productParam := &nmodel.Product{}

	params := mux.Vars(r)
	id := params["id"]

	intID, _ := strconv.Atoi(id)
	uintID := uint(intID)
	err := h.Service.DeleteProduct(uintID)
	if err != nil {
		rf.Response(nrf.ERROR, productParam, w)
		return
	}

	rf.Response(nrf.SUCCESS, productParam, w)
}
