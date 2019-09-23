package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateProduct is business logic for create product
func (s *Service) CreateProduct(product *nmodel.Product) (*nmodel.Product, error) {
	//Send data to repository
	result, err := s.Repository.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return result, nil
}
