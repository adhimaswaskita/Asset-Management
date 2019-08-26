package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//DeleteProductSupplier is used to delete ProductSupplier data from database
func (r *Repository) DeleteProductSupplier(ID uint) error {
	productSupplier := nmodel.ProductSupplier{}
	r.DB.Where("ID = ?", ID).Delete(&productSupplier)

	return nil
}
