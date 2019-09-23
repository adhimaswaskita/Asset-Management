package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetAllManufacture is used to get all manfacture data from database
func (r *Repository) GetAllManufacture() (mManufacture []nmodel.Manufacture, er error) {
	r.DB.Find(&mManufacture)

	return mManufacture, nil
}
