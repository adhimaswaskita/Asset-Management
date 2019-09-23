package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//GetAllAsset is business logic for get all Asset
func (s *Service) GetAllAsset() ([]nmodel.Asset, error) {
	result, err := s.Repository.GetAllAsset()
	if err != nil {
		return nil, err
	}

	return result, nil
}
