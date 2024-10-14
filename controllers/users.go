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

func Listusers(c *gin.Context) {
	var users []models.Users
	initializers.DB.Find(&users)
	c.JSON(200, gin.H{
		"users": users,
	})
}

func Createusers(c *gin.Context) {
	var input models.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user := models.Users{Name: input.Name, Email: input.Email, Age: input.Age}
	initializers.DB.Create(&user)
	c.JSON(200, gin.H{
		"user": user,
	})
}

func Showusers(c *gin.Context) {
	var user models.Users
	initializers.DB.First(&user, c.Param("id"))
	c.JSON(200, gin.H{
		"user": user,
	})
}

func Updateusers(c *gin.Context) {
	var user models.Users
	initializers.DB.First(&user, c.Param("id"))
	var input models.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	initializers.DB.Model(&user).Updates(input)
	c.JSON(200, gin.H{
		"user": user,
	})
}

func Deleteusers(c *gin.Context) {
	initializers.DB.Delete(&models.Users{}, c.Param("id"))
	c.JSON(200, gin.H{
		"message": "deleted",
	})
}
