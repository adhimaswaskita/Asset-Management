package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CountAssetInStore is bussiness logic for get one Asset from database
func (s *Service) CountAssetInStore() (*nmodel.Count, error) {
	total, err := s.Repository.CountAssetInStore()
	if err != nil {
		return nil, err
	}

	return total, nil
}
