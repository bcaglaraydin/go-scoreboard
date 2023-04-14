package main

import (
	"fmt"
	"log"
	"os"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	// database.ConnectDb()
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	SetupRoutes(app)
	// app.Get("/swagger/*", swagger.HandlerDefault)
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
