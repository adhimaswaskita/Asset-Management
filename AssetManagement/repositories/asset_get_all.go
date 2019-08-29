package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetAllAsset is used to get all manfacture data from database
func (r *Repository) GetAllAsset() (Asset []nmodel.Asset, err error) {
	r.DB.Find(&Asset)

	return Asset, nil
}
