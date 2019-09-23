package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateProductType insert data to database from param
func (r *Repository) CreateProductType(mProductType *nmodel.ProductType) (*nmodel.ProductType, error) {
	err := r.DB.Create(&mProductType).Error
	if err != nil {
		return nil, err
	}

	return mProductType, nil
}
