package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//UpdateProductStatus is bussiness logic for product Status update
func (s *Service) UpdateProductStatus(ID uint, ProductStatus *nmodel.ProductStatus) (*nmodel.ProductStatus, error) {

	result, err := s.Repository.UpdateProductStatus(ID, ProductStatus)
	if err != nil {
		return nil, err
	}

	return result, nil
}
