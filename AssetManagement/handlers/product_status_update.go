package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//UpdateProductStatus is used to update product Status
func (h *Handler) UpdateProductStatus(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	productStatusParam := &nmodel.ProductStatus{}

	params := mux.Vars(r)
	ID := params["id"]

	intID, _ := strconv.Atoi(ID)
	uintID := uint(intID)

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(productStatusParam)

	result, err := h.Service.UpdateProductStatus(uintID, productStatusParam)
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, result, nil, w)
}
