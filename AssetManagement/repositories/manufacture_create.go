package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateManufacture is used to insert manufacture data to database
func (r *Repository) CreateManufacture(manufacture *nmodel.Manufacture) (*nmodel.Manufacture, error) {
	err := r.DB.Create(&manufacture).Error
	if err != nil {
		return nil, err
	}

	return manufacture, nil
}
