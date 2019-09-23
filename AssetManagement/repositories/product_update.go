package repositories

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//UpdateProduct is used to save updated data to database
func (r *Repository) UpdateProduct(ID uint, mProduct *nmodel.Product) (*nmodel.Product, error) {
	var mpProduct nmodel.Product

	err := r.DB.Model(&mpProduct).Where("ID = ?", ID).First(&mpProduct).Updates(mProduct).Error
	if err != nil {
		return nil, err
	}
	return mProduct, nil
}
