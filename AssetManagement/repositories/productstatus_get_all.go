package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetAllProductStatus is used to get all manfacture data from database
func (r *Repository) GetAllProductStatus() (productStatus []nmodel.ProductStatus, err error) {
	r.DB.Find(&productStatus)

	return productStatus, nil
}
