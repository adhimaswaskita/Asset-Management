package services

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//UpdateAsset is bussiness logic for Asset  update
func (s *Service) UpdateAsset(ID uint, Asset *nmodel.Asset) (*nmodel.Asset, error) {

	assets, err := s.Repository.UpdateAsset(ID, Asset)
	if err != nil {
		return nil, err
	}

	return assets, nil
}
