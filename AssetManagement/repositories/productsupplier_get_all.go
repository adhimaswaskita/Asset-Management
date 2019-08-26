package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetAllProductSupplier is used to get all manfacture data from database
func (r *Repository) GetAllProductSupplier() (productSupplier []nmodel.ProductSupplier, er error) {
	r.DB.Find(&productSupplier)

	return productSupplier, nil
}
