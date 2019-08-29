package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateAsset is used to insert manufacture data to database
func (r *Repository) CreateAsset(Asset *nmodel.Asset) (*nmodel.Asset, error) {
	err := r.DB.Create(&Asset).Error
	if err != nil {
		return nil, err
	}

	return Asset, nil
}
