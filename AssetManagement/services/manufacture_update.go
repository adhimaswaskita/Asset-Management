package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//UpdateManufacture is business logic for update manufacture
func (s *Service) UpdateManufacture(ID uint, Manufacture *nmodel.Manufacture) (*nmodel.Manufacture, error) {
	manufactures, err := s.Repository.UpdateManufacture(ID, Manufacture)
	if err != nil {
		return nil, err
	}

	return manufactures, nil
}
