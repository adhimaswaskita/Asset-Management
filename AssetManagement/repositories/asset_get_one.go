package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetOneAsset get one Asset detail from Asset table
func (r *Repository) GetOneAsset(ID uint) (*nmodel.Asset, error) {
	var Asset nmodel.Asset
	err := r.DB.Model(&Asset).Where("ID = ?", ID).First(&Asset).Error
	if err != nil {
		return nil, err
	}

	return &Asset, nil
}
