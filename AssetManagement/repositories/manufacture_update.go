package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//UpdateManufacture is used to save updated data to database
func (r *Repository) UpdateManufacture(ID uint, Manufacture *nmodel.Manufacture) (*nmodel.Manufacture, error) {
	var manufactureParam nmodel.Manufacture

	err := r.DB.Model(&manufactureParam).Where("ID = ?", ID).Updates(Manufacture).Error
	if err != nil {
		return nil, err
	}

	return Manufacture, nil
}
