package models

import (
	"github.com/maiga28/guides_gorm/initializers"
	"gorm.io/gorm"
)

// init loads environment variables and initializes the database
func init() {
	initializers.LocalEnvVariables()
	initializers.Database()
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func (Product) TableName() string {
	return "products"
}
