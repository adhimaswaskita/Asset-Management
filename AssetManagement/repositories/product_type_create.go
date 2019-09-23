package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateProductType insert data to database from param
func (r *Repository) CreateProductType(pt *nmodel.ProductType) (*nmodel.ProductType, error) {
	err := r.DB.Create(&pt).Error
	if err != nil {
		return nil, err
	}

	return pt, nil
}
