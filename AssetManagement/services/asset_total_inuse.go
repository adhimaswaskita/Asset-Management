package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//TotalInUseAsset is bussiness logic for get one Asset from database
func (s *Service) TotalInUseAsset() (*nmodel.Dashboard, error) {
	result, err := s.Repository.TotalInUseAsset()
	if err != nil {
		return nil, err
	}

	return result, nil
}
