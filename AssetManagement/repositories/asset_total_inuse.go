package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//TotalInUseAsset get one Asset detail from Asset table
func (r *Repository) TotalInUseAsset() (*nmodel.Dashboard, error) {
	var total int

	err := r.DB.Table("assets").Where("product_status_id = ?", 4).Count(&total).Error
	if err != nil {
		return nil, err
	}

	var result = &nmodel.Dashboard{
		Name:     "In Use",
		Quantity: total,
	}

	return result, nil
}
