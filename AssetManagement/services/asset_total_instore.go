package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//TotalInStoreAsset is bussiness logic for get one Asset from database
func (s *Service) TotalInStoreAsset() (*nmodel.Dashboard, error) {
	result, err := s.Repository.TotalInStoreAsset()
	if err != nil {
		return nil, err
	}

	return result, nil
}
