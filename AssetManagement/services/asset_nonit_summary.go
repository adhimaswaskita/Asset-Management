package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//AssetNonIT is business logic for get all Asset
func (s *Service) AssetNonIT() (*[]nmodel.Dashboard, error) {
	assets, err := s.Repository.AssetNonIT()
	if err != nil {
		return nil, err
	}

	return assets, nil
}
