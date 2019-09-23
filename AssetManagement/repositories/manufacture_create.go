package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateManufacture is used to insert manufacture data to database
func (r *Repository) CreateManufacture(mManufacture *nmodel.Manufacture) (*nmodel.Manufacture, error) {
	err := r.DB.Create(&mManufacture).Error
	if err != nil {
		return nil, err
	}

	return mManufacture, nil
}
