package patient_controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovds/first-go-api/internal/entities"
	patient_services "github.com/joaovds/first-go-api/internal/services/patient"
	"github.com/joaovds/first-go-api/pkg/validation"
)

func GetAllPatients(c *fiber.Ctx) error {
  patients, err := patient_services.GetAllPatients()
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(err)
  }

  return c.Status(fiber.StatusOK).JSON(patients)
}

func GetPatientById(c *fiber.Ctx) error {
  patientId := c.Params("patientId")
  if patientId == "" {
    return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "patientId is required"})
  }

  if !validation.IsValidUUID(patientId) {
    return c.Status(fiber.StatusBadRequest).JSON(
      map[string]string{
        "error": "patientId is not a valid UUID",
      },
    )
  }

  patient, err := patient_services.GetPatientById(patientId)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(err)
  }

  if patient == nil {
    return c.Status(fiber.StatusNoContent).JSON(nil)
  }

  return c.Status(fiber.StatusOK).JSON(patient)
}

func CreatePatient(c *fiber.Ctx) error {
  patient := new(entities.CreatePatientRequest)

  if err := c.BodyParser(patient); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(
      map[string]string{
        "error": "invalid request body",
      },
    )
  }

  createdPatient, err := patient_services.CreatePatient(patient)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(err)
  }

  return c.Status(fiber.StatusCreated).JSON(createdPatient)
}

func UpdatePatient(c *fiber.Ctx) error {
  patientId := c.Params("patientId")
  if patientId == "" {
    return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "patientId is required"})
  }

  if !validation.IsValidUUID(patientId) {
    return c.Status(fiber.StatusBadRequest).JSON(
      map[string]string{
        "error": "patientId is not a valid UUID",
      },
      )
  }

  patient := new(entities.Patient)

  if err := c.BodyParser(patient); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(
      map[string]string{
        "error": "invalid request body",
      },
    )
  }

  err := patient_services.UpdatePatient(patientId, patient)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(err)
  }

  return c.Status(fiber.StatusNoContent).JSON(nil)
}

func DeletePatient(c *fiber.Ctx) error {
  patientId := c.Params("patientId")
  if patientId == "" {
    return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "patientId is required"})
  }

  if !validation.IsValidUUID(patientId) {
    return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "patientId is not a valid UUID"})
  }

  err := patient_services.DeletePatient(patientId)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(err)
  }

  return c.Status(fiber.StatusNoContent).JSON(nil)
}

