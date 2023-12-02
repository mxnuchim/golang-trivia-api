package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mxnuchim/golang-trivia-api/database"
)

func main() {
    // connect to postgreSQL db
    database.ConnectDB()

    app := fiber.New()

    setupRoutes(app)

    log.Fatal(app.Listen(":8080"))
}