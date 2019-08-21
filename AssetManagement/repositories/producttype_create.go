package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateProductType insert data to database from param
func (r *Repository) CreateProductType(pt *nmodel.ProductType) (*nmodel.ProductType, error) {
	r.DB.Create(&pt)
	return pt, nil
}
