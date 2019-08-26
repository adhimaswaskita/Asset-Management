package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetAllProduct is used to get all manfacture data from database
func (r *Repository) GetAllProduct() (product []nmodel.Product, err error) {
	r.DB.Find(&product)

	return product, nil
}
