package handlers

import (
	"net/http"
	"strconv"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	"github.com/gorilla/mux"
)

//DeleteProductSupplier is used to delete product supplier by it's ID
func (h *Handler) DeleteProductSupplier(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	productSupplierParam := &nmodel.ProductSupplier{}

	params := mux.Vars(r)
	id := params["id"]

	intID, _ := strconv.Atoi(id)
	uintID := uint(intID)
	err := h.Service.DeleteProductSupplier(uintID)
	if err != nil {
		rf.Response(nrf.ERROR, productSupplierParam, w)
		return
	}

	rf.Response(nrf.SUCCESS, productSupplierParam, w)
}
