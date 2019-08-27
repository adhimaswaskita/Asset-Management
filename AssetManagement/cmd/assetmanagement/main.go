package main

import (
	"fmt"
	"net/http"

	nconfig "github.com/adhimaswaskita/AssetManagement/config"
	nhandlers "github.com/adhimaswaskita/AssetManagement/handlers"
	nmodels "github.com/adhimaswaskita/AssetManagement/models"
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

	router.HandleFunc("/productsupplier", h.CreateProductSupplier).Methods("POST")
	router.HandleFunc("/productsupplier", h.GetAllProductSupplier).Methods("GET")
	router.HandleFunc("/productsupplier/{id}", h.UpdateProductSupplier).Methods("PUT")
	router.HandleFunc("/productsupplier/{id}", h.DeleteProductSupplier).Methods("DELETE")

	router.HandleFunc("/product", h.CreateProduct).Methods("POST")
	router.HandleFunc("/product", h.GetAllProduct).Methods("GET")
	router.HandleFunc("/product/{id}", h.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/{id}", h.DeleteProduct).Methods("DELETE")

	return router
}

//AutoMigrateDB is database auto migration using gorm
func AutoMigrateDB(repository *nrepo.Repository) {
	repository.DB.AutoMigrate(&nmodels.ProductType{})
	repository.DB.AutoMigrate(&nmodels.Manufacture{})
	repository.DB.AutoMigrate(&nmodels.ProductSupplier{})
	repository.DB.AutoMigrate(&nmodels.Product{})

	repository.DB.Model(&nmodels.Product{}).AddForeignKey("manufacture_id", "manufactures(id)", "RESTRICT", "RESTRICT")
	repository.DB.Model(&nmodels.Product{}).AddForeignKey("product_type_id", "product_types(id)", "RESTRICT", "RESTRICT")
	repository.DB.Model(&nmodels.Product{}).AddForeignKey("product_supplier_id", "product_suppliers(id)", "RESTRICT", "RESTRICT")
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

	AutoMigrateDB(repository)

	service := nservices.NewService(repository)
	handler := nhandlers.NewHandler(service)

	router := NewRouter(handler)

	http.Handle("/", router)

	fmt.Printf("Listening on port%v \n", config.App.Server)
	http.ListenAndServe(config.App.Server, nil)
}
