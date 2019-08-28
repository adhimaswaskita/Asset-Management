package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetAllOrganizationRegion is business logic to get all OrganizationRegion data
func (s *Service) GetAllOrganizationRegion() ([]nmodel.OrganizationRegion, error) {
	organizationRegion, err := s.Repository.GetAllOrganizationRegion()
	if err != nil {
		return nil, err
	}

	return organizationRegion, nil
}
