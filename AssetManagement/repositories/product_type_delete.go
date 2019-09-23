package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//DeleteProductType delete data from database by inserted param
func (r *Repository) DeleteProductType(ID uint) error {
	productType := nmodel.ProductType{}
	err := r.DB.Where("ID = ?", ID).First(&productType).Delete(&productType).Error
	if err != nil {
		return err
	}

	return nil
}
