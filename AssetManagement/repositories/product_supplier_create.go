package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateProductSupplier is used to insert manufacture data to database
func (r *Repository) CreateProductSupplier(mProductSupplier *nmodel.ProductSupplier) (*nmodel.ProductSupplier, error) {
	err := r.DB.Create(&mProductSupplier).Error
	if err != nil {
		return nil, err
	}

	return mProductSupplier, nil
}
