package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//DeleteProduct is used to delete Product data from database
func (r *Repository) DeleteProduct(ID uint) error {
	mProduct := nmodel.Product{}
	err := r.DB.Where("ID = ?", ID).First(&mProduct).Delete(&mProduct).Error
	if err != nil {
		return err
	}

	return nil
}
