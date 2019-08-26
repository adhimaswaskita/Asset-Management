package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//DeleteProduct is used to delete Product data from database
func (r *Repository) DeleteProduct(ID uint) error {
	product := nmodel.Product{}
	r.DB.Where("ID = ?", ID).Delete(&product)

	return nil
}
