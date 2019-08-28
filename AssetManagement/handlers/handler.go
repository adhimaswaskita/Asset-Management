package handlers

import (
	"net/http"

	nservice "github.com/adhimaswaskita/AssetManagement/services"
)

//IHandler is handler contract
type IHandler interface {
	//Product Type
	CreateProductType(w http.ResponseWriter, r *http.Request)
	GetAllProductType(w http.ResponseWriter, r *http.Request)
	UpdateProductType(w http.ResponseWriter, r *http.Request)
	DeleteProductType(w http.ResponseWriter, r *http.Request)
	//Manufacture
	CreateManufacture(w http.ResponseWriter, r *http.Request)
	GetAllManufacture(w http.ResponseWriter, r *http.Request)
	UpdateManufacture(w http.ResponseWriter, r *http.Request)
	DeleteManufacture(w http.ResponseWriter, r *http.Request)
	//ProductSupplier
	CreateProductSupplier(w http.ResponseWriter, r *http.Request)
	GetAllProductSupplier(w http.ResponseWriter, r *http.Request)
	UpdateProductSupplier(w http.ResponseWriter, r *http.Request)
	DeleteProductSupplier(w http.ResponseWriter, r *http.Request)
	//Product
	CreateProduct(w http.ResponseWriter, r *http.Request)
	GetAllProduct(w http.ResponseWriter, r *http.Request)
	GetOneProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
	//Organization
	CreateOrganization(w http.ResponseWriter, r *http.Request)
	GetAllOrganization(w http.ResponseWriter, r *http.Request)
	UpdateOrganization(w http.ResponseWriter, r *http.Request)
	DeleteOrganization(w http.ResponseWriter, r *http.Request)
	//Organization Region
	CreateOrganizationRegion(w http.ResponseWriter, r *http.Request)
	GetAllOrganizationRegion(w http.ResponseWriter, r *http.Request)
	UpdateOrganizationRegion(w http.ResponseWriter, r *http.Request)
	DeleteOrganizationRegion(w http.ResponseWriter, r *http.Request)
}

//Handler is http request handler that implements IHandler
type Handler struct {
	Service nservice.IService
}

//NewHandler creates Handler object
func NewHandler(s nservice.IService) IHandler {
	handler := &Handler{
		Service: s,
	}
	return handler
}
