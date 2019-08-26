package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//UpdateProduct is used to save updated data to database
func (r *Repository) UpdateProduct(ID uint, Product *nmodel.Product) (*nmodel.Product, error) {
	var productParam nmodel.Product

	r.DB.Model(&productParam).Where("ID = ?", ID).Updates(Product)
	return Product, nil
}
