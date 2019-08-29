package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//DeleteProductStatus is used to delete ProductStatus data from database
func (r *Repository) DeleteProductStatus(ID uint) error {
	productStatus := nmodel.ProductStatus{}
	err := r.DB.Where("ID = ?", ID).First(&productStatus).Delete(&productStatus).Error
	if err != nil {
		return err
	}

	return nil
}
