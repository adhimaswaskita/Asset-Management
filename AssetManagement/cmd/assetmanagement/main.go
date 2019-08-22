package main

import (
	"fmt"
	"net/http"

	nconfig "github.com/adhimaswaskita/AssetManagement/config"
	nhandlers "github.com/adhimaswaskita/AssetManagement/handlers"
	nrepo "github.com/adhimaswaskita/AssetManagement/repositories"
	nservices "github.com/adhimaswaskita/AssetManagement/services"
	"github.com/gorilla/mux"
)

const (
	//FilePath is config filepath
	FilePath = "config.yaml"
)

//NewRouter creates new mux.Router
func NewRouter(h nhandlers.IHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/producttype", h.CreateProductType).Methods("POST")
	router.HandleFunc("/producttype", h.GetAllProductType).Methods("GET")
	router.HandleFunc("/producttype/{id}", h.UpdateProductType).Methods("PUT")
	router.HandleFunc("/producttype/{id}", h.DeleteProductType).Methods("DELETE")

	router.HandleFunc("/manufacture", h.CreateManufacture).Methods("POST")
	router.HandleFunc("/manufacture", h.GetAllManufacture).Methods("GET")
	router.HandleFunc("/manufacture/{id}", h.UpdateManufacture).Methods("PUT")
	router.HandleFunc("/manufacture/{id}", h.DeleteManufacture).Methods("DELETE")
	return router
}

func main() {

	config, err := nconfig.NewConfig(FilePath)
	if err != nil {
		fmt.Printf("Error read config file : %v", err)
	}

	repository, err := nrepo.NewRepository(config.Repo)
	if err != nil {
		fmt.Printf("Failed connect to database : %v", err)
		return
	}

	service := nservices.NewService(repository)
	handler := nhandlers.NewHandler(service)

	router := NewRouter(handler)

	http.Handle("/", router)

	fmt.Printf("Listening on port%v \n", config.App.Server)
	http.ListenAndServe(config.App.Server, nil)
}
