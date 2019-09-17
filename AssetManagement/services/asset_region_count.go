package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//AssetByRegion is business logic for get all Asset by it's region
func (s *Service) AssetByRegion() (*[]nmodel.Dashboard, error) {
	assets, err := s.Repository.AssetByRegion()
	if err != nil {
		return nil, err
	}

	return assets, nil
}
