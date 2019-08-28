package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//UpdateOrganization is business logic for update Organization
func (s *Service) UpdateOrganization(ID uint, Organization *nmodel.Organization) (*nmodel.Organization, error) {
	organizations, err := s.Repository.UpdateOrganization(ID, Organization)
	if err != nil {
		return nil, err
	}

	return organizations, nil
}
