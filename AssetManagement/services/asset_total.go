package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//TotalAsset is bussiness logic for get one Asset from database
func (s *Service) TotalAsset() (*nmodel.Dashboard, error) {
	result, err := s.Repository.TotalAsset()
	if err != nil {
		return nil, err
	}

	return result, nil
}
