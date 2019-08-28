package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateOrganizationRegion is business logic for Create OrganizationRegion
func (s *Service) CreateOrganizationRegion(m *nmodel.OrganizationRegion) (*nmodel.OrganizationRegion, error) {
	organizationRegion, err := s.Repository.CreateOrganizationRegion(m)
	if err != nil {
		return nil, err
	}

	return organizationRegion, nil
}
