package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CountAssetInRepair get one Asset detail from Asset table
func (r *Repository) CountAssetInRepair() (*nmodel.Count, error) {
	var total int

	err := r.DB.Table("assets").Where("product_status_id = ?", 3).Count(&total).Error
	if err != nil {
		return nil, err
	}

	var count = &nmodel.Count{
		Name:     "In Repair",
		Quantity: total,
	}

	return count, nil
}
