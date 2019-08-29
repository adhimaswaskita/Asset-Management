package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//DeleteAsset is used to delete Asset data from database
func (r *Repository) DeleteAsset(ID uint) error {
	Asset := nmodel.Asset{}
	err := r.DB.Where("ID = ?", ID).First(&Asset).Delete(&Asset).Error
	if err != nil {
		return err
	}

	return nil
}
