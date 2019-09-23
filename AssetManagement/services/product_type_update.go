package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//UpdateProductType is bussiness logic for product type update
func (s *Service) UpdateProductType(ID uint, ProductType *nmodel.ProductType) (*nmodel.ProductType, error) {

	result, err := s.Repository.UpdateProductType(ID, ProductType)
	if err != nil {
		return nil, err
	}

	return result, nil
}
