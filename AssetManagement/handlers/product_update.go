package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//UpdateProduct is used to update product
func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	productParam := &nmodel.Product{}

	params := mux.Vars(r)
	id := params["id"]

	intID, _ := strconv.Atoi(id)
	ID := uint(intID)

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(productParam)

	product, err := h.Service.UpdateProduct(ID, productParam)
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, product, nil, w)
}
