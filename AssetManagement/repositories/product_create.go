package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateProduct is used to insert manufacture data to database
func (r *Repository) CreateProduct(product *nmodel.Product) (*nmodel.Product, error) {
	r.DB.Create(&product)

	return product, nil
}
