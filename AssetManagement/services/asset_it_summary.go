package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//AssetIT is business logic for get all Asset
func (s *Service) AssetIT() (*[]nmodel.Dashboard, error) {
	assets, err := s.Repository.AssetIT()
	if err != nil {
		return nil, err
	}

	return assets, nil
}
