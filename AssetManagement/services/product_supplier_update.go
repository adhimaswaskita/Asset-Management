package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//UpdateProductSupplier is bussiness logic for product supplier update
func (s *Service) UpdateProductSupplier(ID uint, ProductSupplier *nmodel.ProductSupplier) (*nmodel.ProductSupplier, error) {

	result, err := s.Repository.UpdateProductSupplier(ID, ProductSupplier)
	if err != nil {
		return nil, err
	}

	return result, nil
}
