package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//UpdateProductSupplier is used to save updated data to database
func (r *Repository) UpdateProductSupplier(ID uint, ProductSupplier *nmodel.ProductSupplier) (*nmodel.ProductSupplier, error) {
	var productSupplierParam nmodel.ProductSupplier

	r.DB.Model(&productSupplierParam).Where("ID = ?", ID).Updates(ProductSupplier)
	return ProductSupplier, nil
}
