package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//GetAllProductStatus is business logic for get all product Status
func (s *Service) GetAllProductStatus() ([]nmodel.ProductStatus, error) {
	productStatuses, err := s.Repository.GetAllProductStatus()
	if err != nil {
		return nil, err
	}

	return productStatuses, nil
}
