package handlers

import (
	"net/http"
	"strconv"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	"github.com/gorilla/mux"
)

//DeleteProductType is used to delete product type by it's ID
func (h *Handler) DeleteProductType(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}
	productTypeParam := &nmodel.ProductType{}

	params := mux.Vars(r)
	id := params["id"]

	intID, _ := strconv.Atoi(id)
	err := h.Service.DeleteProductType(intID)
	if err != nil {
		rf.Response(nrf.ERROR, productTypeParam, w)
		return
	}

	rf.Response(nrf.SUCCESS, productTypeParam, w)
}
