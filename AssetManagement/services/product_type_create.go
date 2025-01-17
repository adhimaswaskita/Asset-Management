package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateProductType is business logic for create product type
func (s *Service) CreateProductType(pt *nmodel.ProductType) (*nmodel.ProductType, error) {
	//Send data to repository
	result, err := s.Repository.CreateProductType(pt)
	if err != nil {
		return nil, err
	}

	return result, nil
}
