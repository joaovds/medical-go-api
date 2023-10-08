package patient_controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovds/first-go-api/internal/services"
)

func GetAllPatients(c *fiber.Ctx) error {
  patients, err := patient_services.GetAllPatients()
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(err)
  }

  return c.JSON(patients)
}

