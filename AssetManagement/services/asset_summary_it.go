package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//ITAssetSummary is business logic for get all Asset
func (s *Service) ITAssetSummary() (*[]nmodel.Dashboard, error) {
	assets, err := s.Repository.ITAssetSummary()
	if err != nil {
		return nil, err
	}

	return assets, nil
}
