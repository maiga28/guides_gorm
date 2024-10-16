package initializers

import (
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() {
	// Récupérer la variable d'environnement
	dsn := os.Getenv("DB_URLS")
	if dsn == "" {
		log.Fatal("DB_URLS environment variable is not set")
		return
	}

	// Se connecter à la base de données
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established")
}
