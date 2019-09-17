package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//AssetNonITOther is business logic for get all Asset
func (s *Service) AssetNonITOther() (*nmodel.Dashboard, error) {
	quantityContainer := 0

	assets, err := s.Repository.AssetNonIT()
	if err != nil {
		return nil, err
	}

	for i, asset := range *assets {
		if i >= 5 {
			quantityContainer += asset.Quantity
		}
	}

	OtherAsset := nmodel.Dashboard{
		Name:     "Others",
		Quantity: quantityContainer,
	}

	return &OtherAsset, nil
}
