package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetAllOrganization is business logic to get all Organization data
func (s *Service) GetAllOrganization() ([]nmodel.Organization, error) {
	result, err := s.Repository.GetAllOrganization()
	if err != nil {
		return nil, err
	}

	return result, nil
}
