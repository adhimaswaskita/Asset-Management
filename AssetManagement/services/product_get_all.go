package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//GetAllProduct is business logic for get all product
func (s *Service) GetAllProduct() ([]nmodel.Product, error) {
	products, err := s.Repository.GetAllProduct()
	if err != nil {
		return nil, err
	}

	return products, nil
}
