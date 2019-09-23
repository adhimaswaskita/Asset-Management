package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateProductStatus is used to insert manufacture data to database
func (r *Repository) CreateProductStatus(mProductStatus *nmodel.ProductStatus) (*nmodel.ProductStatus, error) {
	err := r.DB.Create(&mProductStatus).Error
	if err != nil {
		return nil, err
	}

	return mProductStatus, nil
}
