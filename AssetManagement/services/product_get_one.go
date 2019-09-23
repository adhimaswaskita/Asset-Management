package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetOneProduct is bussiness logic for get one product from database
func (s *Service) GetOneProduct(ID uint) (*nmodel.Product, error) {
	result, err := s.Repository.GetOneProduct(ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
