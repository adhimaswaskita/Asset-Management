package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CountAsset is bussiness logic for get one Asset from database
func (s *Service) CountAsset() (*nmodel.Count, error) {
	total, err := s.Repository.CountAsset()
	if err != nil {
		return nil, err
	}

	return total, nil
}
