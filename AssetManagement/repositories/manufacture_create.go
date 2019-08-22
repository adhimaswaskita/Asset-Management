package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateManufacture is used to insert manufacture data to database
func (r *Repository) CreateManufacture(manufacture *nmodel.Manufacture) (*nmodel.Manufacture, error) {
	r.DB.Create(&manufacture)

	return manufacture, nil
}
