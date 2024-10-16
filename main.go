// package main

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/maiga28/guides_gorm/controllers"
// 	"github.com/maiga28/guides_gorm/initializers"
// 	"log"
// )

// // init loads environment variables and initializes the database
// func init() {
// 	initializers.LocalEnvVariables()
// 	initializers.Database()
// }

package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/maiga28/guides_gorm/controllers"
	"github.com/maiga28/guides_gorm/initializers"
)

func init() {
	initializers.LocalEnvVariables()
	initializers.Database()
}

func main() {
	r := gin.Default()

	// Middleware pour ignorer favicon.ico
	r.Use(func(c *gin.Context) {
		if c.Request.RequestURI == "/favicon.ico" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Afficher toutes les routes au démarrage
	for _, route := range r.Routes() {
		log.Printf("Route enregistrée : %s %s\n", route.Method, route.Path)
	}

	// Groupes de routes pour les utilisateurs
	users := r.Group("/users")
	{
		users.GET("", controllers.Listusers)          // /users
		users.GET("/:id", controllers.Showusers)      // /users/:id
		users.POST("", controllers.Createusers)       // /users
		users.PUT("/:id", controllers.Updateusers)    // /users/:id
		users.DELETE("/:id", controllers.Deleteusers) // /users/:id
	}

	// Groupes de routes pour les produits
	products := r.Group("/products")
	{
		products.GET("", controllers.ProductsIndex)        // /products
		products.GET("/:id", controllers.ShowProduct)      // /products/:id
		products.POST("", controllers.CreateProduct)       // /products
		products.PUT("/:id", controllers.UpdateProduct)    // /products/:id
		products.DELETE("/:id", controllers.DeleteProduct) // /products/:id
	}

	// Lancer le serveur
	r.Run()
}
