package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//TopFiveITAssets is business logic for get all Asset
func (s *Service) TopFiveITAssets() (*[]nmodel.Dashboard, error) {
	var assetResult []nmodel.Dashboard
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
			assetResult = append(assetResult, asset)
		} else {
			OthersAsset.Quantity += asset.Quantity
		}
	}

	assetResult = append(assetResult, OthersAsset)

	return &assetResult, nil
}
