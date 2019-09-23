package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetOneAsset get one Asset detail from Asset table
func (r *Repository) GetOneAsset(ID uint) (*nmodel.Asset, error) {
	var mAsset nmodel.Asset
	err := r.DB.Model(&mAsset).Where("ID = ?", ID).First(&mAsset).Error
	if err != nil {
		return nil, err
	}

	return &mAsset, nil
}
