package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//DeleteProduct is used to delete Product data from database
func (r *Repository) DeleteProduct(ID uint) error {
	product := nmodel.Product{}
	err := r.DB.Where("ID = ?", ID).First(&product).Delete(&product).Error
	if err != nil {
		return err
	}

	return nil
}
