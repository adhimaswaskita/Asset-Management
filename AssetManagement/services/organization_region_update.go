package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//UpdateOrganizationRegion is business logic for update OrganizationRegion
func (s *Service) UpdateOrganizationRegion(ID uint, OrganizationRegion *nmodel.OrganizationRegion) (*nmodel.OrganizationRegion, error) {
	organizationRegions, err := s.Repository.UpdateOrganizationRegion(ID, OrganizationRegion)
	if err != nil {
		return nil, err
	}

	return organizationRegions, nil
}
