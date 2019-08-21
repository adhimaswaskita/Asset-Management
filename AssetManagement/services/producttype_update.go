package services

import nmodel "github.com/adhimaswaskita/netmonk-asset-management/models"

//UpdateProductType is bussiness logic for product type update
func (s *Service) UpdateProductType(ID uint, ProductType *nmodel.ProductType) (*nmodel.ProductType, error) {

	productTypes, err := s.Repository.UpdateProductType(ID, ProductType)
	if err != nil {
		return nil, err
	}

	return productTypes, nil
}
