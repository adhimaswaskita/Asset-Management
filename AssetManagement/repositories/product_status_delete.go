package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//DeleteProductStatus is used to delete ProductStatus data from database
func (r *Repository) DeleteProductStatus(ID uint) error {
	mProductStatus := nmodel.ProductStatus{}
	err := r.DB.Where("ID = ?", ID).First(&mProductStatus).Delete(&mProductStatus).Error
	if err != nil {
		return err
	}

	return nil
}
