package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//DeleteProductSupplier is used to delete ProductSupplier data from database
func (r *Repository) DeleteProductSupplier(ID uint) error {
	mProductSupplier := nmodel.ProductSupplier{}
	err := r.DB.Where("ID = ?", ID).First(&mProductSupplier).Delete(&mProductSupplier).Error
	if err != nil {
		return err
	}

	return nil
}
