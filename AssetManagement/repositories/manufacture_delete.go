package repositories

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//DeleteManufacture is used to delete manufacture data from database
func (r *Repository) DeleteManufacture(ID uint) error {
	mManufacture := nmodel.Manufacture{}
	err := r.DB.Where("ID = ?", ID).First(&mManufacture).Delete(&mManufacture).Error

	if err != nil {
		return err
	}

	return nil
}
