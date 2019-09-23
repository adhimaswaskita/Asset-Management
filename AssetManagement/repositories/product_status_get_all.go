package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetAllProductStatus is used to get all manfacture data from database
func (r *Repository) GetAllProductStatus() (mProductStatus []nmodel.ProductStatus, err error) {
	r.DB.Find(&mProductStatus)

	return mProductStatus, nil
}
