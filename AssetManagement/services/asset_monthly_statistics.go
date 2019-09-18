package services

import (
	"strconv"

	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//AssetStatistics is business logic for get all Asset
func (s *Service) AssetStatistics(year int) (*[]nmodel.Statistics, error) {
	var AssetResult []nmodel.Statistics
	assets, err := s.Repository.AssetStatistics(year)

	months := [12]string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DES"}
	if err != nil {
		return nil, err
	}

	for _, asset := range *assets {
		intMonth, err := strconv.Atoi(asset.Month)
		if err != nil {
			return nil, err
		}
		asset.Month = months[intMonth-1]
		AssetResult = append(AssetResult, asset)
	}

	return &AssetResult, nil
}
