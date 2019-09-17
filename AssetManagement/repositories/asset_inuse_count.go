package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CountAssetInUse get one Asset detail from Asset table
func (r *Repository) CountAssetInUse() (*nmodel.StatusCount, error) {
	var total int

	err := r.DB.Table("assets").Where("product_status_id = ?", 4).Count(&total).Error
	if err != nil {
		return nil, err
	}

	var count = &nmodel.StatusCount{
		Name:     "In Use",
		Quantity: total,
	}

	return count, nil
}
