package consultation_controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovds/first-go-api/internal/entities"
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

func CreateConsultation(c *fiber.Ctx) error {
  consultation := new(entities.CreateConsultationRequest)

  if err := c.BodyParser(consultation); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(
      map[string]string{
        "error": "invalid request body",
      },
    )
  }

  if validationError := consultation.Validate(); validationError.Message != "" {
    return c.Status(fiber.StatusBadRequest).JSON(validationError)
  }

  err := consultation_services.CreateConsultation(consultation)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
  }

  return c.Status(fiber.StatusCreated).JSON(nil)
}

func DeleteConsultation(c *fiber.Ctx) error {
  consultationId := c.Params("consultationId")
  if consultationId == "" {
    return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "consultationId is required"})
  }

  if !validation.IsValidUUID(consultationId) {
    return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "consultationId is not a valid UUID"})
  }

  err := consultation_services.DeleteConsultation(consultationId)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(err)
  }

  return c.Status(fiber.StatusNoContent).JSON(nil)
}

