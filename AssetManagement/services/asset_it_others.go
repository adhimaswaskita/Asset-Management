package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//AssetITOther is business logic for get all Asset
func (s *Service) AssetITOther() (*nmodel.AssetSummary, error) {
	quantityContainer := 0

	assets, err := s.Repository.AssetIT()
	if err != nil {
		return nil, err
	}

	for i, asset := range *assets {
		if i >= 5 {
			quantityContainer += asset.Quantity
		}
	}

	OtherAsset := nmodel.AssetSummary{
		Name:     "Others",
		Quantity: quantityContainer,
	}

	return &OtherAsset, nil
}
