package configs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovds/first-go-api/internal/infra/fiber/middlewares"

  "github.com/joaovds/first-go-api/internal/infra/fiber/routes"
)

func GetApp() *fiber.App {
  app := fiber.New()
  
  middlewares.SetupMiddlewares(app)
  routes.SetupRoutes(app)

  return app
}

