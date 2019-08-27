package repositories

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//DeleteManufacture is used to delete manufacture data from database
func (r *Repository) DeleteManufacture(ID uint) error {
	manufacture := nmodel.Manufacture{}
	err := r.DB.Where("ID = ?", ID).First(&manufacture).Delete(&manufacture).Error

	if err != nil {
		return err
	}

	return nil
}
