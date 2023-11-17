package entities

import (
	"fmt"
	"time"

	"github.com/joaovds/first-go-api/pkg/validation"
)

type Consultation struct {
	Id          string    `json:"id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Notes       string    `json:"notes"`
	Diagnosis   string    `json:"diagnosis"`
	Doctor      Doctor    `json:"doctor"`
	Patient     Patient   `json:"patient"`
}

type CreateConsultationRequest struct {
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
  Notes       string    `json:"notes"`
  Diagnosis   string    `json:"diagnosis"`
  DoctorId    string    `json:"doctorId"`
  PatientId   string    `json:"patientId"`
}

func (c *CreateConsultationRequest) Validate() ValidationError {
  if c.Date.IsZero() {
    return ValidationError{Message: "date is required"}
  }

  if c.Date.Before(time.Now()) {
    fmt.Println(time.Now())
    return ValidationError{Message: "date must be a valid date and greater than today"}
  }

  if c.DoctorId == "" {
    return ValidationError{Message: "doctorId is required"}
  }

  if !validation.IsValidUUID(c.DoctorId) {
    return ValidationError{Message: "doctorId is not a valid UUID"}
  }

  if c.PatientId == "" {
    return ValidationError{Message: "patientId is required"}
  }

  if !validation.IsValidUUID(c.PatientId) {
    return ValidationError{Message: "patientId is not a valid UUID"}
  }

  return ValidationError{}
}

