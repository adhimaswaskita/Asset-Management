package repositories

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//UpdateAsset is used to save updated data to database
func (r *Repository) UpdateAsset(ID uint, Asset *nmodel.Asset) (*nmodel.Asset, error) {
	var assetParam nmodel.Asset

	err := r.DB.Model(&assetParam).Where("ID = ?", ID).First(&assetParam).Updates(Asset).Error
	if err != nil {
		return nil, err
	}
	return Asset, nil
}
