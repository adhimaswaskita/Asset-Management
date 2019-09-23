package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetOneProduct get one product detail from product table
func (r *Repository) GetOneProduct(ID uint) (*nmodel.Product, error) {
	var mProduct nmodel.Product
	err := r.DB.Model(&mProduct).Where("ID = ?", ID).First(&mProduct).Error
	if err != nil {
		return nil, err
	}

	return &mProduct, nil
}
