package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//UpdateProduct is bussiness logic for product  update
func (s *Service) UpdateProduct(ID uint, Product *nmodel.Product) (*nmodel.Product, error) {

	result, err := s.Repository.UpdateProduct(ID, Product)
	if err != nil {
		return nil, err
	}

	return result, nil
}
