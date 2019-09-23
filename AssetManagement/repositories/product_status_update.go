package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//UpdateProductStatus is used to save updated data to database
func (r *Repository) UpdateProductStatus(ID uint, ProductStatus *nmodel.ProductStatus) (*nmodel.ProductStatus, error) {
	var productStatusParam nmodel.ProductStatus

	err := r.DB.Model(&productStatusParam).Where("ID = ?", ID).First(&productStatusParam).Updates(ProductStatus).Error
	if err != nil {
		return nil, err
	}

	return ProductStatus, nil
}
