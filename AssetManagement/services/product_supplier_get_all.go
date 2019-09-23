package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//GetAllProductSupplier is business logic for get all product supplier
func (s *Service) GetAllProductSupplier() ([]nmodel.ProductSupplier, error) {
	result, err := s.Repository.GetAllProductSupplier()
	if err != nil {
		return nil, err
	}

	return result, nil
}
