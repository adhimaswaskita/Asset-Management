package handlers

import (
	"net/http"
	"strconv"

	nrf "github.com/adhimaswaskita/AssetManagement/lib/responseformat"
	"github.com/gorilla/mux"
)

//AssetStatistics is used to get all Asset  data
func (h *Handler) AssetStatistics(w http.ResponseWriter, r *http.Request) {
	rf := nrf.ResponseFormat{}

	params := mux.Vars(r)
	yearString := params["year"]

	year, _ := strconv.Atoi(yearString)

	result, err := h.Service.AssetStatistics(year)
	if err != nil {
		stringErr := err.Error()
		rf.Response(nrf.ERROR, nil, stringErr, w)
		return
	}

	rf.Response(nrf.SUCCESS, result, nil, w)
	return
}
