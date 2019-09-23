package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateManufacture is business logic for Create Manufacture
func (s *Service) CreateManufacture(m *nmodel.Manufacture) (*nmodel.Manufacture, error) {
	result, err := s.Repository.CreateManufacture(m)
	if err != nil {
		return nil, err
	}

	return result, nil
}
