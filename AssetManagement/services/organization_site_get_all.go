package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetAllOrganizationSite is business logic to get all OrganizationSite data
func (s *Service) GetAllOrganizationSite() ([]nmodel.OrganizationSite, error) {
	result, err := s.Repository.GetAllOrganizationSite()
	if err != nil {
		return nil, err
	}

	return result, nil
}
