package routes

import (
  "github.com/gofiber/fiber/v2"
)

func handlePatientRoutes(router fiber.Router) {
  patientRouter := router.Group("/patient")

  patientRouter.Get("/", func(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
      "hello": "world!",
    })
  })
}


