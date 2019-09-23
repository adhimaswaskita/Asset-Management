package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateProductSupplier is business logic for create product supplier
func (s *Service) CreateProductSupplier(productSupplier *nmodel.ProductSupplier) (*nmodel.ProductSupplier, error) {
	//Send data to repository
	result, err := s.Repository.CreateProductSupplier(productSupplier)
	if err != nil {
		return nil, err
	}

	return result, nil
}
