package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//GetAllAsset is business logic for get all Asset
func (s *Service) GetAllAsset() ([]nmodel.Asset, error) {
	assets, err := s.Repository.GetAllAsset()
	if err != nil {
		return nil, err
	}

	return assets, nil
}
