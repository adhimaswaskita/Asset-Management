package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetOneProduct get one product detail from product table
func (r *Repository) GetOneProduct(ID uint) (*nmodel.Product, error) {
	var product nmodel.Product
	err := r.DB.Model(&product).Where("ID = ?", ID).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}
