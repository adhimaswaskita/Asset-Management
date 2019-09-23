package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//TotalInRepairAsset is bussiness logic for get one Asset from database
func (s *Service) TotalInRepairAsset() (*nmodel.Dashboard, error) {
	result, err := s.Repository.TotalInRepairAsset()
	if err != nil {
		return nil, err
	}

	return result, nil
}
