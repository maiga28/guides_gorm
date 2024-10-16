package models

import (
	"github.com/maiga28/guides_gorm/initializers"
)

// init loads environment variables and initializes the database
func init() {
	initializers.LocalEnvVariables()
	initializers.Database()
}

type Product struct {
	ID    uint `gorm:"primaryKey"` // GORM manages primary keys automatically
	Code  string
	Price uint
}

func (Product) TableName() string {
	return "products"
}
