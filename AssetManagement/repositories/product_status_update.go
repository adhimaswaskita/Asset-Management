package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//UpdateProductStatus is used to save updated data to database
func (r *Repository) UpdateProductStatus(ID uint, mProductStatus *nmodel.ProductStatus) (*nmodel.ProductStatus, error) {
	var mpProductStatus nmodel.ProductStatus

	err := r.DB.Model(&mpProductStatus).Where("ID = ?", ID).First(&mpProductStatus).Updates(mProductStatus).Error
	if err != nil {
		return nil, err
	}

	return mProductStatus, nil
}
