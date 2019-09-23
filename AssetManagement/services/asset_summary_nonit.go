package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//NonITAssetSummary is business logic for get all Asset
func (s *Service) NonITAssetSummary() (*[]nmodel.Dashboard, error) {
	assets, err := s.Repository.NonITAssetSummary()
	if err != nil {
		return nil, err
	}

	return assets, nil
}
