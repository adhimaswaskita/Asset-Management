package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetAllProduct is used to get all manfacture data from database
func (r *Repository) GetAllProduct() (mProduct []nmodel.Product, err error) {
	r.DB.Find(&mProduct)

	return mProduct, nil
}
