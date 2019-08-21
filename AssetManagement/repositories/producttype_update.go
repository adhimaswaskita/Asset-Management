package repositories

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//UpdateProductType ...
func (r *Repository) UpdateProductType(ID uint, productType *nmodel.ProductType) (*nmodel.ProductType, error) {
	var pt nmodel.ProductType

	r.DB.Model(&pt).Where("ID = ?", ID).Updates(productType)

	return productType, nil
}
