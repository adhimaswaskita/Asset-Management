package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//DeleteProductType delete data from database by inserted param
func (r *Repository) DeleteProductType(ID uint) error {
	mProductType := nmodel.ProductType{}
	err := r.DB.Where("ID = ?", ID).First(&mProductType).Delete(&mProductType).Error
	if err != nil {
		return err
	}

	return nil
}
