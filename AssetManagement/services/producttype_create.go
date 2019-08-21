package services

import nmodel "github.com/adhimaswaskita/netmonk-asset-management/models"

//CreateProductType is business logic for create product type
func (s *Service) CreateProductType(pt *nmodel.ProductType) (*nmodel.ProductType, error) {
	//Send data to repository
	productTypes, err := s.Repository.CreateProductType(pt)
	if err != nil {
		return nil, err
	}

	return productTypes, nil
}
