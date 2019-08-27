package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateProductSupplier is used to insert manufacture data to database
func (r *Repository) CreateProductSupplier(productSupplier *nmodel.ProductSupplier) (*nmodel.ProductSupplier, error) {
	err := r.DB.Create(&productSupplier).Error
	if err != nil {
		return nil, err
	}

	return productSupplier, nil
}
