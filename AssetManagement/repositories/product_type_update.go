package repositories

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//UpdateProductType ...
func (r *Repository) UpdateProductType(ID uint, mProductType *nmodel.ProductType) (*nmodel.ProductType, error) {
	var mpProductType nmodel.ProductType

	err := r.DB.Model(&mpProductType).Where("ID = ?", ID).First(&mpProductType).Updates(mProductType).Error
	if err != nil {
		return nil, err
	}

	return mProductType, nil
}
