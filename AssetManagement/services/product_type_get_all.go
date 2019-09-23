package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//GetAllProductType is business logic for get all product type
func (s *Service) GetAllProductType() ([]nmodel.ProductType, error) {
	result, err := s.Repository.GetAllProductType()
	if err != nil {
		return nil, err
	}

	return result, nil
}
