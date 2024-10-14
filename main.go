package main

import (
	// "log"
	// "github.com/gin-gonic/gin"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/maiga28/guides_gorm/controllers"
	"github.com/maiga28/guides_gorm/initializers"
)

// init loads environment variables and initializes the database
func init() {
	initializers.LocalEnvVariables()
	initializers.Database()
}

func main() {
	r := gin.Default()
	// Middleware pour afficher les routes
	r.Use(func(c *gin.Context) {
		c.Next() // Continue avec le traitement de la requête
		// Affiche toutes les routes à chaque requête
		for _, route := range r.Routes() {
			log.Printf("%s %s\n", route.Method, route.Path)
		}
	})

	r.GET("/liste/users", controllers.Listusers)
	r.GET("/users/:id", controllers.Showusers)
	r.POST("/create/users", controllers.Createusers)
	r.PUT("/update/users/:id", controllers.Updateusers)
	r.DELETE("/delete/users/:id", controllers.Deleteusers)

	r.GET("/liste/products", controllers.ProductsIndex)
	r.GET("/products/:id", controllers.Show)
	r.POST("/create/products", controllers.Create)
	r.PUT("/update/products/:id", controllers.Update)
	r.DELETE("/delete/products/:id", controllers.Delete)

	// Lancer le serveur
	r.Run()
}
