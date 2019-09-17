package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CountAssetInUse is bussiness logic for get one Asset from database
func (s *Service) CountAssetInUse() (*nmodel.StatusCount, error) {
	total, err := s.Repository.CountAssetInUse()
	if err != nil {
		return nil, err
	}

	return total, nil
}
