package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//AssetBySite is business logic for get all Asset by it's region
func (s *Service) AssetBySite() (*[]nmodel.Dashboard, error) {
	assets, err := s.Repository.AssetBySite()
	if err != nil {
		return nil, err
	}

	return assets, nil
}
