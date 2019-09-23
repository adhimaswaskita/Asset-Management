package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateProduct is used to insert manufacture data to database
func (r *Repository) CreateProduct(mProduct *nmodel.Product) (*nmodel.Product, error) {
	err := r.DB.Create(&mProduct).Error
	if err != nil {
		return nil, err
	}

	return mProduct, nil
}
