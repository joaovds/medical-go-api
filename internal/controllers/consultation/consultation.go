package consultation_controllers

import (
  "github.com/gofiber/fiber/v2"
  consultation_services "github.com/joaovds/first-go-api/internal/services/consultation"
)

func GetAllConsultations(c *fiber.Ctx) error {
  consultations, err := consultation_services.GetAllConsultations()
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(err)
  }

  return c.Status(fiber.StatusOK).JSON(consultations)
}

