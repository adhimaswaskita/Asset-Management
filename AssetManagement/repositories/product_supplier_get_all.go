package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetAllProductSupplier is used to get all manfacture data from database
func (r *Repository) GetAllProductSupplier() (mProductSupplier []nmodel.ProductSupplier, err error) {
	r.DB.Find(&mProductSupplier)

	return mProductSupplier, nil
}
