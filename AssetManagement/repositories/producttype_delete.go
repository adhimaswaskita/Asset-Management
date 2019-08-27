package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//DeleteProductType delete data from database by inserted param
func (r *Repository) DeleteProductType(ID int) error {
	productType := nmodel.ProductType{}
	err := r.DB.Where("ID = ?", ID).Delete(&productType).Error
	if err != nil {
		return err
	}

	return nil
}
