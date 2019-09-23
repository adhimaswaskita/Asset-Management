package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//DeleteAsset is used to delete Asset data from database
func (r *Repository) DeleteAsset(ID uint) error {
	mAsset := nmodel.Asset{}
	err := r.DB.Where("ID = ?", ID).First(&mAsset).Delete(&mAsset).Error
	if err != nil {
		return err
	}

	return nil
}
