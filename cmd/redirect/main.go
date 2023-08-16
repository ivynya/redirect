package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	log.Println("[redirect] creating router...")

	createRouter(app)

	tokenPresent := os.Getenv("NOTION_TOKEN") != ""
	log.Println("[redirect] $NOTION_DB_ID:", os.Getenv("NOTION_DB_ID"))
	log.Println("[redirect] $NOTION_TOKEN:", tokenPresent)
	log.Println("[redirect] $ANALYTICS_HOST:", os.Getenv("ANALYTICS_HOST"))
	log.Println("[redirect] $ANALYTICS_VERSION:", os.Getenv("ANALYTICS_VERSION"))
	log.Println("[redirect] starting fiber on port 3000")

	log.Fatal(app.Listen(":3000"))
}
