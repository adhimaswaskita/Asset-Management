package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//TotalInUseAsset get one Asset detail from Asset table
func (r *Repository) TotalInUseAsset() (*nmodel.Dashboard, error) {
	var assetTotal int

	err := r.DB.Table("assets").Where("product_status_id = ?", 4).Count(&assetTotal).Error
	if err != nil {
		return nil, err
	}

	var result = &nmodel.Dashboard{
		Name:     "In Use",
		Quantity: assetTotal,
	}

	return result, nil
}
