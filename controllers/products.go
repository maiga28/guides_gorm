package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/maiga28/guides_gorm/initializers"
	"github.com/maiga28/guides_gorm/models"
	// "gorm.io/gorm"
	"net/http"
)

func init() {
	initializers.LocalEnvVariables()
	initializers.Database()
}

// List all products
func ProductsIndex(c *gin.Context) {
	var products []models.Product
	initializers.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

// Create a new product (accepts both JSON and form data)
func CreateProduct(c *gin.Context) {
	var input models.Product
	// Try to bind the request body as JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		// If JSON binding fails, attempt to bind form data
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	// Create a new product using the input data
	product := models.Product{
		ID:    0,
		Code:  input.Code,
		Price: input.Price,
	}
	initializers.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{
		"message": "Produit ajouté avec succès",
		"product": product,
	})
}

// Show a specific product by ID
func ShowProduct(c *gin.Context) {
	var product models.Product
	if err := initializers.DB.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produit non trouvé"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

// Update a product by ID
func UpdateProduct(c *gin.Context) {
	var product models.Product
	if err := initializers.DB.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produit non trouvé"})
		return
	}

	// Bind the input data
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the product
	initializers.DB.Model(&product).Updates(input)
	c.JSON(http.StatusOK, gin.H{
		"message": "Produit mis à jour avec succès",
		"product": product,
	})
}

// Delete a product by ID
func DeleteProduct(c *gin.Context) {
	if err := initializers.DB.Delete(&models.Product{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Échec de la suppression du produit"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Produit supprimé avec succès",
	})
}
