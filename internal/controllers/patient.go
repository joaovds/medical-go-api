package patient_controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovds/first-go-api/internal/services"
	"github.com/joaovds/first-go-api/pkg/validation"
)

func GetAllPatients(c *fiber.Ctx) error {
  patients, err := patient_services.GetAllPatients()
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(err)
  }

  return c.JSON(patients)
}

func GetPatientById(c *fiber.Ctx) error {
  patientId := c.Params("patientId")
  if patientId == "" {
    return c.Status(fiber.StatusBadRequest).JSON("patientId is required")
  }

  if !validation.IsValidUUID(patientId) {
    return c.Status(fiber.StatusBadRequest).JSON("patientId is not a valid UUID")
  }

  patient, err := patient_services.GetPatientById(patientId)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(err)
  }

  return c.JSON(patient)
}

