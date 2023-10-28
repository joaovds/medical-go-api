package middlewares

import (
  "github.com/gofiber/fiber/v2"
)

func SetupMiddlewares(app *fiber.App) {
  setupCors(app)
  setupLogger(app)
}
