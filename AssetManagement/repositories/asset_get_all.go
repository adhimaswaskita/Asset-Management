package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetAllAsset is used to get all manfacture data from database
func (r *Repository) GetAllAsset() (mAsset []nmodel.Asset, err error) {
	r.DB.Find(&mAsset)

	return mAsset, nil
}
