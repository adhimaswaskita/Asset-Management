package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetAllManufacture is business logic to get all manufacture data
func (s *Service) GetAllManufacture() ([]nmodel.Manufacture, error) {
	manufactures, err := s.Repository.GetAllManufacture()
	if err != nil {
		return nil, err
	}

	return manufactures, nil
}
