package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//TotalAssetByRegion is business logic for get all Asset by it's region
func (s *Service) TotalAssetByRegion() (*[]nmodel.Dashboard, error) {
	result, err := s.Repository.TotalAssetByRegion()
	if err != nil {
		return nil, err
	}

	return result, nil
}
