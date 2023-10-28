package middlewares

import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"
)

func setupCors(app *fiber.App) {
  app.Use(cors.New(cors.Config{
    AllowOrigins: "*",
    AllowHeaders: "Origin, Content-Type, Accept",
  }))
}
