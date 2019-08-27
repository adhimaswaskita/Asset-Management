package repositories

import (
	"fmt"
	"strings"

	nconfig "github.com/adhimaswaskita/AssetManagement/config"
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
	"github.com/jinzhu/gorm"

	//Import postgres from gorm and initialize it
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//IRepository is repo contract
type IRepository interface {
	//Product Type
	GetAllProductType() ([]nmodel.ProductType, error)
	CreateProductType(*nmodel.ProductType) (*nmodel.ProductType, error)
	UpdateProductType(ID uint, ProductType *nmodel.ProductType) (*nmodel.ProductType, error)
	DeleteProductType(ID int) error
	//Manufacture
	CreateManufacture(*nmodel.Manufacture) (*nmodel.Manufacture, error)
	GetAllManufacture() ([]nmodel.Manufacture, error)
	UpdateManufacture(ID uint, Manufacture *nmodel.Manufacture) (*nmodel.Manufacture, error)
	DeleteManufacture(ID uint) error
	//Product Supplier
	GetAllProductSupplier() ([]nmodel.ProductSupplier, error)
	CreateProductSupplier(*nmodel.ProductSupplier) (*nmodel.ProductSupplier, error)
	UpdateProductSupplier(ID uint, ProductSupplier *nmodel.ProductSupplier) (*nmodel.ProductSupplier, error)
	DeleteProductSupplier(ID uint) error
	//Product
	GetAllProduct() ([]nmodel.Product, error)
	GetOneProduct(ID uint) (*nmodel.Product, error)
	CreateProduct(*nmodel.Product) (*nmodel.Product, error)
	UpdateProduct(ID uint, Product *nmodel.Product) (*nmodel.Product, error)
	DeleteProduct(ID uint) error
}

//Repository implements IRepository
type Repository struct {
	DB *gorm.DB
}

//NewRepository is used to open connection to postgres database by config passes from parameter
func NewRepository(repoConfig *nconfig.Repository) (*Repository, error) {
	server := strings.Split(repoConfig.Server, ":")
	postgresConfig := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v",
		server[0], server[1], repoConfig.Username, repoConfig.Name, repoConfig.Password)

	db, err := gorm.Open("postgres", postgresConfig)
	if err != nil {
		return nil, err
	}
	// defer db.Close()

	repository := &Repository{
		DB: db,
	}

	return repository, nil
}
