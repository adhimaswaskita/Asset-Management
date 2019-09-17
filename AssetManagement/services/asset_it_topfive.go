package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//AssetITTopFive is business logic for get all Asset
func (s *Service) AssetITTopFive() (*[]nmodel.Dashboard, error) {
	var LimitedAsset []nmodel.Dashboard
	assets, err := s.Repository.AssetIT()
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
