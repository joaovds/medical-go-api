package entities

import (
	"net/mail"
	"time"
)

type Patient struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Document    string `json:"document"`
	DateOfBirth string `json:"dateOfBirth"`
}

type CreatePatientRequest struct {
	Patient
}

func (c *CreatePatientRequest) Validate() ValidationError {
  if c.Name == "" {
    return ValidationError{Message: "Name is required"}
  }

  if c.Email == "" {
    return ValidationError{Message: "Email is required"}
  }
  _, err := mail.ParseAddress(c.Email)
  if err != nil {
    return ValidationError{Message: "Email is invalid"}
  }

  if c.Document == "" {
    return ValidationError{Message: "Document is required"}
  }

  if c.DateOfBirth == "" {
    return ValidationError{Message: "Date of birth is required"}
  }
  parseDate, err := time.Parse("2006-01-02", c.DateOfBirth);
  if err != nil {
    return ValidationError{Message: "Date of birth is invalid, must be in the format YYYY-MM-DD"}
  }
  if time.Now().Sub(parseDate) < 0 {
    return ValidationError{Message: "Date of birth is invalid, can't be in the future"}
  }

  return ValidationError{}
}
