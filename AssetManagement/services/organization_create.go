package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateOrganization is business logic for Create Organization
func (s *Service) CreateOrganization(m *nmodel.Organization) (*nmodel.Organization, error) {
	result, err := s.Repository.CreateOrganization(m)
	if err != nil {
		return nil, err
	}

	return result, nil
}
