package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateOrganizationSite is business logic for Create OrganizationSite
func (s *Service) CreateOrganizationSite(m *nmodel.OrganizationSite) (*nmodel.OrganizationSite, error) {
	organizationSite, err := s.Repository.CreateOrganizationSite(m)
	if err != nil {
		return nil, err
	}

	return organizationSite, nil
}