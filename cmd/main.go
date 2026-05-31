package main

import (
    "log"
    "os"

    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
    
    "github.com/suphachok09790/url-shortener/internal/database"
)

func main() {
    godotenv.Load()
    
    database.Connect()
    
    app := fiber.New()

    app.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"status": "ok"})
    })

    port := os.Getenv("APP_PORT")
    log.Fatal(app.Listen(":" + port))
}