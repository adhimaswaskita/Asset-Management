package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//DeleteProductSupplier is used to delete ProductSupplier data from database
func (r *Repository) DeleteProductSupplier(ID uint) error {
	productSupplier := nmodel.ProductSupplier{}
	err := r.DB.Where("ID = ?", ID).First(&productSupplier).Delete(&productSupplier).Error
	if err != nil {
		return err
	}

	return nil
}
