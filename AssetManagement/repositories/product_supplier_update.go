package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//UpdateProductSupplier is used to save updated data to database
func (r *Repository) UpdateProductSupplier(ID uint, mProductSupplier *nmodel.ProductSupplier) (*nmodel.ProductSupplier, error) {
	var mpProductSupplier nmodel.ProductSupplier

	err := r.DB.Model(&mpProductSupplier).Where("ID = ?", ID).First(&mpProductSupplier).Updates(mProductSupplier).Error
	if err != nil {
		return nil, err
	}

	return mProductSupplier, nil
}
