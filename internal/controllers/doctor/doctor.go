package doctor_controllers

import (
	"github.com/gofiber/fiber/v2"
	doctor_services "github.com/joaovds/first-go-api/internal/services/doctor"
	"github.com/joaovds/first-go-api/pkg/validation"
)

func GetAllDoctors(c *fiber.Ctx) error {
  doctors, err := doctor_services.GetAllDoctors()
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(err)
  }

  return c.Status(fiber.StatusOK).JSON(doctors)
}

func GetDoctorById(c *fiber.Ctx) error {
  doctorId := c.Params("doctorId")
  if doctorId == "" {
    return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "doctorId is required"})
  }

  if !validation.IsValidUUID(doctorId) {
    return c.Status(fiber.StatusBadRequest).JSON(
      map[string]string{
        "error": "doctorId is not a valid UUID",
      },
    )
  }

  doctor, err := doctor_services.GetDoctorById(doctorId)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(err)
  }

  if doctor == nil {
    return c.Status(fiber.StatusNoContent).JSON(nil)
  }

  return c.Status(fiber.StatusOK).JSON(doctor)
}
