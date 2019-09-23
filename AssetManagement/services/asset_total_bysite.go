package services

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//TotalAssetBySite is business logic for get all Asset by it's region
func (s *Service) TotalAssetBySite() (*[]nmodel.Dashboard, error) {
	result, err := s.Repository.TotalAssetBySite()
	if err != nil {
		return nil, err
	}

	return result, nil
}
