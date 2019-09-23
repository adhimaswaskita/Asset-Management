package repositories

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//UpdateAsset is used to save updated data to database
func (r *Repository) UpdateAsset(ID uint, mAsset *nmodel.Asset) (*nmodel.Asset, error) {
	var mpAsset nmodel.Asset

	err := r.DB.Model(&mpAsset).Where("ID = ?", ID).First(&mpAsset).Updates(mAsset).Error
	if err != nil {
		return nil, err
	}
	return mAsset, nil
}
