package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//TotalInStoreAsset get one Asset detail from Asset table
func (r *Repository) TotalInStoreAsset() (*nmodel.Dashboard, error) {
	var assetTotal int

	err := r.DB.Table("assets").Where("product_status_id = ?", 2).Count(&assetTotal).Error
	if err != nil {
		return nil, err
	}

	var result = &nmodel.Dashboard{
		Name:     "In Store",
		Quantity: assetTotal,
	}

	return result, nil
}
