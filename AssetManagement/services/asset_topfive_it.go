package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//TopFiveITAssets is business logic for get all Asset
func (s *Service) TopFiveITAssets() (*[]nmodel.Dashboard, error) {
	var result []nmodel.Dashboard
	assets, err := s.Repository.ITAssetSummary()
	if err != nil {
		return nil, err
	}

	OthersAsset := nmodel.Dashboard{
		Name:     "Others",
		Quantity: 0,
	}
	for i, asset := range *assets {
		if i < 5 {
			result = append(result, asset)
		} else {
			OthersAsset.Quantity += asset.Quantity
		}
	}

	result = append(result, OthersAsset)

	return &result, nil
}
