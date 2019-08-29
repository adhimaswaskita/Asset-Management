package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateProductStatus is business logic for create product Status
func (s *Service) CreateProductStatus(productStatus *nmodel.ProductStatus) (*nmodel.ProductStatus, error) {
	//Send data to repository
	productStatuses, err := s.Repository.CreateProductStatus(productStatus)
	if err != nil {
		return nil, err
	}

	return productStatuses, nil
}
