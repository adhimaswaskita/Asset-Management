package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//UpdateOrganizationSite is business logic for update OrganizationSite
func (s *Service) UpdateOrganizationSite(ID uint, OrganizationSite *nmodel.OrganizationSite) (*nmodel.OrganizationSite, error) {
	result, err := s.Repository.UpdateOrganizationSite(ID, OrganizationSite)
	if err != nil {
		return nil, err
	}

	return result, nil
}
