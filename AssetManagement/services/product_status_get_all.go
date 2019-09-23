package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//GetAllProductStatus is business logic for get all product Status
func (s *Service) GetAllProductStatus() ([]nmodel.ProductStatus, error) {
	result, err := s.Repository.GetAllProductStatus()
	if err != nil {
		return nil, err
	}

	return result, nil
}
