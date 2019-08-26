package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//UpdateProductSupplier is used to update product Supplier
func (h *Handler) UpdateProductSupplier(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	productSupplierParam := &nmodel.ProductSupplier{}

	params := mux.Vars(r)
	id := params["id"]

	intID, _ := strconv.Atoi(id)
	ID := uint(intID)

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(productSupplierParam)

	productSupplier, err := h.Service.UpdateProductSupplier(ID, productSupplierParam)
	if err != nil {
		rf.Response(nrf.ERROR, productSupplier, w)
		return
	}

	rf.Response(nrf.SUCCESS, productSupplier, w)
}
