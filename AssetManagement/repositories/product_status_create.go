package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateProductStatus is used to insert manufacture data to database
func (r *Repository) CreateProductStatus(productStatus *nmodel.ProductStatus) (*nmodel.ProductStatus, error) {
	err := r.DB.Create(&productStatus).Error
	if err != nil {
		return nil, err
	}

	return productStatus, nil
}
