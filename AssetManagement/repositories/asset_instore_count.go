package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CountAssetInStore get one Asset detail from Asset table
func (r *Repository) CountAssetInStore() (*nmodel.Dashboard, error) {
	var total int

	err := r.DB.Table("assets").Where("product_status_id = ?", 2).Count(&total).Error
	if err != nil {
		return nil, err
	}

	var count = &nmodel.Dashboard{
		Name:     "In Store",
		Quantity: total,
	}

	return count, nil
}
