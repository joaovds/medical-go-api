package configs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovds/first-go-api/internal/configs/middlewares"

  "github.com/joaovds/first-go-api/internal/configs/routes"
)

func GetApp() *fiber.App {
  app := fiber.New()
  
  middlewares.SetupMiddlewares(app)
  routes.SetupRoutes(app)

  return app
}

