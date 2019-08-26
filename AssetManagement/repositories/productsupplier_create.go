package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateProductSupplier is used to insert manufacture data to database
func (r *Repository) CreateProductSupplier(productSupplier *nmodel.ProductSupplier) (*nmodel.ProductSupplier, error) {
	r.DB.Create(&productSupplier)

	return productSupplier, nil
}
