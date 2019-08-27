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
	DeleteProductType(ID int) error
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
