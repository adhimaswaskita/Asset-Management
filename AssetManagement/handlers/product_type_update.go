package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//UpdateProductType is used to update product type
func (h *Handler) UpdateProductType(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	productTypeParam := &nmodel.ProductType{}

	params := mux.Vars(r)
	ID := params["id"]

	intID, _ := strconv.Atoi(ID)
	uintID := uint(intID)

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(productTypeParam)

	result, err := h.Service.UpdateProductType(uintID, productTypeParam)
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, result, nil, w)
}
