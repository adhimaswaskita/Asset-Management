package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CountAssetInUse get one Asset detail from Asset table
func (r *Repository) CountAssetInUse() (*nmodel.Count, error) {
	var total int

	err := r.DB.Table("assets").Where("product_status_id = ?", 4).Count(&total).Error
	if err != nil {
		return nil, err
	}

	var count = &nmodel.Count{
		Name:     "In Use",
		Quantity: total,
	}

	return count, nil
}
