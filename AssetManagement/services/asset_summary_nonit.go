package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//NonITAssetSummary is business logic for get all Asset
func (s *Service) NonITAssetSummary() (*[]nmodel.Dashboard, error) {
	result, err := s.Repository.NonITAssetSummary()
	if err != nil {
		return nil, err
	}

	return result, nil
}
