package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetOneAsset is bussiness logic for get one Asset from database
func (s *Service) GetOneAsset(ID uint) (*nmodel.Asset, error) {
	Asset, err := s.Repository.GetOneAsset(ID)
	if err != nil {
		return nil, err
	}

	return Asset, nil
}
