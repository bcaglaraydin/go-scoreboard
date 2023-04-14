package main

import (
	"fmt"
	"log"
	"os"

	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/bcaglaraydin/go-scoreboard/docs"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

//	@title			Go Scoreboard API
//	@version		1.0
//	@description	Ranking users based on their scores

//	@contact.name	Berdan Çağlar AYDIN
//	@contact.url	linkedin.com/in/bcaglaraydin/
//	@contact.email	berdancaglaraydin@gmail.com

//	@host		localhost:3000
//	@BasePath	/

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
	app.Get("/swagger/*", swagger.HandlerDefault)
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
