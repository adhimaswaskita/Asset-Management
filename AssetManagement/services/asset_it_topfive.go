package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//AssetITTopFive is business logic for get all Asset
func (s *Service) AssetITTopFive() (*[]nmodel.Dashboard, error) {
	var ResultAsset []nmodel.Dashboard
	assets, err := s.Repository.AssetIT()
	if err != nil {
		return nil, err
	}

	OthersAsset := nmodel.Dashboard{
		Name:     "Others",
		Quantity: 0,
	}
	for i, asset := range *assets {
		if i < 5 {
			ResultAsset = append(ResultAsset, asset)
		} else {
			OthersAsset.Quantity += asset.Quantity
		}
	}

	ResultAsset = append(ResultAsset, OthersAsset)

	return &ResultAsset, nil
}
