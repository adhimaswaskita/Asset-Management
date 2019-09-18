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
	//Organization Site
	CreateOrganizationSite(w http.ResponseWriter, r *http.Request)
	GetAllOrganizationSite(w http.ResponseWriter, r *http.Request)
	UpdateOrganizationSite(w http.ResponseWriter, r *http.Request)
	DeleteOrganizationSite(w http.ResponseWriter, r *http.Request)
	//ProductStatus
	CreateProductStatus(w http.ResponseWriter, r *http.Request)
	GetAllProductStatus(w http.ResponseWriter, r *http.Request)
	UpdateProductStatus(w http.ResponseWriter, r *http.Request)
	DeleteProductStatus(w http.ResponseWriter, r *http.Request)
	//Asset
	CreateAsset(w http.ResponseWriter, r *http.Request)
	GetAllAsset(w http.ResponseWriter, r *http.Request)
	GetOneAsset(w http.ResponseWriter, r *http.Request)
	UpdateAsset(w http.ResponseWriter, r *http.Request)
	DeleteAsset(w http.ResponseWriter, r *http.Request)
	CountAsset(w http.ResponseWriter, r *http.Request)
	CountAssetInStore(w http.ResponseWriter, r *http.Request)
	CountAssetInUse(w http.ResponseWriter, r *http.Request)
	CountAssetInRepair(w http.ResponseWriter, r *http.Request)
	AssetIT(w http.ResponseWriter, r *http.Request)
	AssetITTopFive(w http.ResponseWriter, r *http.Request)
	AssetNonIT(w http.ResponseWriter, r *http.Request)
	AssetNonITTopFive(w http.ResponseWriter, r *http.Request)
	AssetByRegion(w http.ResponseWriter, r *http.Request)
	AssetBySite(w http.ResponseWriter, r *http.Request)
	AssetStatistics(w http.ResponseWriter, r *http.Request)
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
