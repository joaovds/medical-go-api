package consultation_controllers

import (
	"github.com/gofiber/fiber/v2"
	consultation_services "github.com/joaovds/first-go-api/internal/services/consultation"
	"github.com/joaovds/first-go-api/pkg/validation"
)

func GetAllConsultations(c *fiber.Ctx) error {
  consultations, err := consultation_services.GetAllConsultations()
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(err)
  }

  return c.Status(fiber.StatusOK).JSON(consultations)
}

func GetConsultationById(c *fiber.Ctx) error {
  consultationId := c.Params("consultationId")
  if consultationId == "" {
    return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "consultationId is required"})
  }

  if !validation.IsValidUUID(consultationId) {
    return c.Status(fiber.StatusBadRequest).JSON(
      map[string]string{
        "error": "consultationId is not a valid UUID",
      },
      )
  }

  consultation, err := consultation_services.GetConsultationById(consultationId)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(err)
  }

  if consultation == nil {
    return c.Status(fiber.StatusNoContent).JSON(nil)
  }

  return c.Status(fiber.StatusOK).JSON(consultation)
}

