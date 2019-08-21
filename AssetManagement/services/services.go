package services

import (
	nmodel "github.com/adhimaswaskita/netmonk-asset-management/models"
	nrepo "github.com/adhimaswaskita/netmonk-asset-management/repositories"
)

//IService is service contract
type IService interface {
	CreateProductType(*nmodel.ProductType) (*nmodel.ProductType, error)
	GetAllProductType() ([]nmodel.ProductType, error)
	UpdateProductType(ID uint, ProductType *nmodel.ProductType) (*nmodel.ProductType, error)
	DeleteProductType(ID int) error
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
