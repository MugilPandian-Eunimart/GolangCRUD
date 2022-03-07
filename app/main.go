package main

import (
	"fiber-mongo-api/utils"
	"fiber-mongo-api/routes"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
  
    //run database
    utils.ConnectDB()

	//routes
    routes.UserRoute(app)
  
    app.Listen(":6000")
}