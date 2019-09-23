package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetAllManufacture is business logic to get all Manufacture data
func (s *Service) GetAllManufacture() ([]nmodel.Manufacture, error) {
	result, err := s.Repository.GetAllManufacture()
	if err != nil {
		return nil, err
	}

	return result, nil
}
