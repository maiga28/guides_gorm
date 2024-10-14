package main

import (
	"github.com/maiga28/guides_gorm/initializers"
	"github.com/maiga28/guides_gorm/models"
)

func init() {
	initializers.LocalEnvVariables()
	initializers.Database()
}
func main() {
	// Migrate the schema
	initializers.DB.AutoMigrate(&models.Product{})
	initializers.DB.AutoMigrate(&models.Users{})
}
