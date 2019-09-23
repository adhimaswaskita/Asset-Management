package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateAsset is business logic for create Asset
func (s *Service) CreateAsset(Asset *nmodel.Asset) (*nmodel.Asset, error) {
	//Send data to repository
	result, err := s.Repository.CreateAsset(Asset)
	if err != nil {
		return nil, err
	}

	return result, nil
}
