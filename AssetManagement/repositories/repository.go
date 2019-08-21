package repositories

import (
	"fmt"
	"strings"

	nconfig "github.com/adhimaswaskita/netmonk-asset-management/config"
	nmodel "github.com/adhimaswaskita/netmonk-asset-management/models"
	"github.com/jinzhu/gorm"

	//Import postgres from gorm and initialize it
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//IRepository is repo contract
type IRepository interface {
	GetAllProductType() ([]nmodel.ProductType, error)
	CreateProductType(*nmodel.ProductType) (*nmodel.ProductType, error)
	UpdateProductType(ID uint, ProductType *nmodel.ProductType) (*nmodel.ProductType, error)
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
