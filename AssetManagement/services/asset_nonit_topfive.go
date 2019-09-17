package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//AssetNonITTopFive is business logic for get all Asset
func (s *Service) AssetNonITTopFive() (*[]nmodel.AssetSummary, error) {
	var LimitedAsset []nmodel.AssetSummary
	assets, err := s.Repository.AssetNonIT()
	if err != nil {
		return nil, err
	}

	for i, asset := range *assets {
		if i < 5 {
			LimitedAsset = append(LimitedAsset, asset)
		}
	}

	return &LimitedAsset, nil
}
