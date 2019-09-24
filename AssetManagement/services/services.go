package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	nrepo "github.com/adhimaswaskita/AssetManagement/repositories"
)

//IService is service contract
type IService interface {
	//Product Type
	CreateProductType(*nmodel.ProductType) (*nmodel.ProductType, error)
	GetAllProductType() ([]nmodel.ProductType, error)
	UpdateProductType(ID uint, ProductType *nmodel.ProductType) (*nmodel.ProductType, error)
	DeleteProductType(ID uint) error

	//Manufacture
	CreateManufacture(*nmodel.Manufacture) (*nmodel.Manufacture, error)
	GetAllManufacture() ([]nmodel.Manufacture, error)
	UpdateManufacture(ID uint, Manufacture *nmodel.Manufacture) (*nmodel.Manufacture, error)
	DeleteManufacture(ID uint) error

	//Product Supplier
	GetAllProductSupplier() ([]nmodel.ProductSupplier, error)
	CreateProductSupplier(*nmodel.ProductSupplier) (*nmodel.ProductSupplier, error)
	UpdateProductSupplier(ID uint, ProductSupplier *nmodel.ProductSupplier) (*nmodel.ProductSupplier, error)
	DeleteProductSupplier(ID uint) error

	//Product
	GetAllProduct() ([]nmodel.Product, error)
	GetOneProduct(ID uint) (*nmodel.Product, error)
	CreateProduct(*nmodel.Product) (*nmodel.Product, error)
	UpdateProduct(ID uint, Product *nmodel.Product) (*nmodel.Product, error)
	DeleteProduct(ID uint) error

	//Organization
	GetAllOrganization() ([]nmodel.Organization, error)
	CreateOrganization(*nmodel.Organization) (*nmodel.Organization, error)
	UpdateOrganization(ID uint, Organization *nmodel.Organization) (*nmodel.Organization, error)
	DeleteOrganization(ID uint) error

	//OrganizationRegion
	GetAllOrganizationRegion() ([]nmodel.OrganizationRegion, error)
	CreateOrganizationRegion(*nmodel.OrganizationRegion) (*nmodel.OrganizationRegion, error)
	UpdateOrganizationRegion(ID uint, OrganizationRegion *nmodel.OrganizationRegion) (*nmodel.OrganizationRegion, error)
	DeleteOrganizationRegion(ID uint) error

	//OrganizationSite
	GetAllOrganizationSite() ([]nmodel.OrganizationSite, error)
	CreateOrganizationSite(*nmodel.OrganizationSite) (*nmodel.OrganizationSite, error)
	UpdateOrganizationSite(ID uint, OrganizationSite *nmodel.OrganizationSite) (*nmodel.OrganizationSite, error)
	DeleteOrganizationSite(ID uint) error

	//Product Supplier
	GetAllProductStatus() ([]nmodel.ProductStatus, error)
	CreateProductStatus(*nmodel.ProductStatus) (*nmodel.ProductStatus, error)
	UpdateProductStatus(ID uint, ProductStatus *nmodel.ProductStatus) (*nmodel.ProductStatus, error)
	DeleteProductStatus(ID uint) error

	//Asset
	GetAllAsset() ([]nmodel.Asset, error)
	GetOneAsset(ID uint) (*nmodel.Asset, error)
	CreateAsset(*nmodel.Asset) (*nmodel.Asset, error)
	UpdateAsset(ID uint, Asset *nmodel.Asset) (*nmodel.Asset, error)
	DeleteAsset(ID uint) error

	//Dashboard
	TotalAsset() (*nmodel.Dashboard, error)
	TotalInStoreAsset() (*nmodel.Dashboard, error)
	TotalInUseAsset() (*nmodel.Dashboard, error)
	TotalInRepairAsset() (*nmodel.Dashboard, error)
	ITAssetSummary() (*[]nmodel.Dashboard, error)
	TopFiveITAssets() (*[]nmodel.Dashboard, error)
	NonITAssetSummary() (*[]nmodel.Dashboard, error)
	TopFiveNonITAssets() (*[]nmodel.Dashboard, error)
	TotalAssetByRegion() (*[]nmodel.Dashboard, error)
	TotalAssetBySite() (*[]nmodel.Dashboard, error)
	AssetStatistics(year int) (*[]nmodel.Statistics, error)
	AssetRecentInsertion() (*[]nmodel.RecentInsertion, error)
}

//Service is business logic that implements IService
type Service struct {
	Repository nrepo.IRepository
}

//NewService creates service object
func NewService(r nrepo.IRepository) IService {
	service := &Service{
		Repository: r,
	}
	return service
}
