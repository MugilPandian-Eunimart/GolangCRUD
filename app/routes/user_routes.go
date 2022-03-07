package routes

import (
    "fiber-mongo-api/services"
    "github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
    app.Post("/user", services.CreateUser)
    app.Get("/user/:userId", services.GetAUser)
    app.Put("/user/:userId", services.EditAUser)
    app.Delete("/user/:userId", services.DeleteAUser)
    app.Get("/users", services.GetAllUsers)
}