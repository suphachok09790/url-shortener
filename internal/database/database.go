package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/suphachok09790/url-shortener/internal/models"
)

var DB *gorm.DB

func Connect() {
    dsn := os.Getenv("DATABASE_URL")
    
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    db.AutoMigrate(&models.URL{})
    
    DB = db
    log.Println("Database connected successfully")
}