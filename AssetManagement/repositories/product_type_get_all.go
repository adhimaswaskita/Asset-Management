package repositories

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//GetAllProductType get all product type from database
func (r *Repository) GetAllProductType() (mProductType []nmodel.ProductType, err error) {
	r.DB.Find(&mProductType)

	return mProductType, nil
}
