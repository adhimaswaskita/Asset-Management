package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//TotalInRepairAsset get one Asset detail from Asset table
func (r *Repository) TotalInRepairAsset() (*nmodel.Dashboard, error) {
	var assetTotal int

	err := r.DB.Table("assets").Where("product_status_id = ?", 3).Count(&assetTotal).Error
	if err != nil {
		return nil, err
	}

	var result = &nmodel.Dashboard{
		Name:     "In Repair",
		Quantity: assetTotal,
	}

	return result, nil
}
