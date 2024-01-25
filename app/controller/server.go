package main

import (

	"github.com/gofiber/fiber/v2"
	"log"
)


func main() {
    app := fiber.New()

	app.Static("/", "../Views/index.html")

    log.Fatal(app.Listen(":3000"))
}