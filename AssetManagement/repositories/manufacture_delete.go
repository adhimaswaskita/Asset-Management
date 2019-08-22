package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//DeleteManufacture is used to delete manufacture data from database
func (r *Repository) DeleteManufacture(ID uint) error {
	manufacture := nmodel.Manufacture{}
	r.DB.Where("ID = ?", ID).Delete(&manufacture)

	return nil
}
