package controllers

import (
	// "log"
	// "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"github.com/maiga28/guides_gorm/initializers"
	"github.com/maiga28/guides_gorm/models"
)

func init() {
	initializers.LocalEnvVariables()
	initializers.Database()
}

func ProductsIndex(c *gin.Context) {
	var products []models.Product
	initializers.DB.Find(&products)
	c.JSON(200, gin.H{
		"products": products,
	})
}

func Create(c *gin.Context) {
	// Validate input
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Create product
	product := models.Product{Code: input.Code, Price: input.Price}
	initializers.DB.Create(&product)

	c.JSON(200, gin.H{
		"product": product,
	})

}

func Show(c *gin.Context) {
	var product models.Product
	initializers.DB.First(&product, c.Param("id"))
	c.JSON(200, gin.H{
		"product": product,
	})
}

func Update(c *gin.Context) {
	// Get model if exist
	var product models.Product
	initializers.DB.First(&product, c.Param("id"))
	// Validate input
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Update product
	initializers.DB.Model(&product).Updates(input)
	c.JSON(200, gin.H{
		"product": product,
	})
}

func Delete(c *gin.Context) {
	initializers.DB.Delete(&models.Product{}, c.Param("id"))
	c.JSON(200, gin.H{
		"message": "deleted",
	})
}
