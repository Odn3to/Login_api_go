package main

import (
	"fmt"
	"log"

    "github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"Login-api/router"
	"Login-api/database"
)

func main() {

	//carregar as variáveis de ambiente
	if err := godotenv.Load("/home/utm/login-api/.env"); err != nil {
        log.Print("No .env file found")
    }

	
	database.ConectaNoBD()
	
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	router.Register(app)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", "8006")))
}