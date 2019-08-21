package repositories

import (
	nmodel "github.com/adhimaswaskita/netmonk-asset-management/models"
)

//GetAllProductType get all product type from database
func (r *Repository) GetAllProductType() (p []nmodel.ProductType, err error) {
	r.DB.Find(&p)

	return p, nil
}
