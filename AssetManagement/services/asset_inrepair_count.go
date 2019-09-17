package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CountAssetInRepair is bussiness logic for get one Asset from database
func (s *Service) CountAssetInRepair() (*nmodel.Dashboard, error) {
	total, err := s.Repository.CountAssetInRepair()
	if err != nil {
		return nil, err
	}

	return total, nil
}
