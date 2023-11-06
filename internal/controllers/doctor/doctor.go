package doctor_controllers

import (
  "github.com/gofiber/fiber/v2"
  doctor_services "github.com/joaovds/first-go-api/internal/services/doctor"
)

func GetAllDoctors(c *fiber.Ctx) error {
  doctors, err := doctor_services.GetAllDoctors()
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(err)
  }

  return c.Status(fiber.StatusOK).JSON(doctors)
}

