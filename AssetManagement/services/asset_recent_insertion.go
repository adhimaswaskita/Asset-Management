package services

import (
	"strconv"
	"strings"

	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//AssetRecentInsertion is business logic for get all Asset
func (s *Service) AssetRecentInsertion() (*[]nmodel.RecentInsertion, error) {
	var result []nmodel.RecentInsertion
	recentAssets, err := s.Repository.AssetRecentInsertion()

	months := [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "Desember"}
	if err != nil {
		return nil, err
	}

	for _, asset := range *recentAssets {
		splitResult := strings.Split(asset.Time, "/")
		monthParam, err := strconv.Atoi(splitResult[1])
		if err != nil {
			return nil, err
		}

		month := months[monthParam-1]

		monthConvert := strings.Replace(splitResult[1], splitResult[1], month, -1)

		splitResult[1] = monthConvert

		timeConcatination := (splitResult[0] + " " + splitResult[1] + " " + splitResult[2])

		asset.Time = timeConcatination

		result = append(result, asset)
	}

	return &result, nil
}
