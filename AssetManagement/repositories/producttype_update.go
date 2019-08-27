package repositories

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//UpdateProductType ...
func (r *Repository) UpdateProductType(ID uint, productType *nmodel.ProductType) (*nmodel.ProductType, error) {
	var productTypeParam nmodel.ProductType

	err := r.DB.Model(&productTypeParam).Where("ID = ?", ID).First(&productTypeParam).Updates(productType).Error
	if err != nil {
		return nil, err
	}

	return productType, nil
}
