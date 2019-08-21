package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	nrf "github.com/adhimaswaskita/netmonk-asset-management/lib/responseformat"
	nmodel "github.com/adhimaswaskita/netmonk-asset-management/models"
)

//UpdateProductType is used to update product type
func (h *Handler) UpdateProductType(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	productTypeParam := &nmodel.ProductType{}

	params := mux.Vars(r)
	id := params["id"]

	intID, _ := strconv.Atoi(id)
	ID := uint(intID)

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(productTypeParam)

	productType, err := h.Service.UpdateProductType(ID, productTypeParam)
	if err != nil {
		rf.Response(nrf.ERROR, productType, w)
		return
	}

	rf.Response(nrf.SUCCESS, productType, w)
}
