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
