package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateAsset is used to insert manufacture data to database
func (r *Repository) CreateAsset(mAsset *nmodel.Asset) (*nmodel.Asset, error) {
	err := r.DB.Create(&mAsset).Error
	if err != nil {
		return nil, err
	}

	return mAsset, nil
}
