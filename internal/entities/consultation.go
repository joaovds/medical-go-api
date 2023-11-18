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
	Date        string    `json:"date"`
	Description string    `json:"description"`
  Notes       string    `json:"notes"`
  Diagnosis   string    `json:"diagnosis"`
  DoctorId    string    `json:"doctorId"`
  PatientId   string    `json:"patientId"`
}

func (c *CreateConsultationRequest) Validate() ValidationError {
  if c.Date == "" {
    return ValidationError{Message: "date is required"}
  }

  parseDate, err := time.Parse("2006-01-02 15:04:05", c.Date)
  if err != nil {
    return ValidationError{Message: "date must be in the format YYYY-MM-DD HH:MM:SS"}
  }

  if parseDate.Before(time.Now()) {
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

