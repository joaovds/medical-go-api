package entities

import "time"

type Patient struct {
  Id       string `json:"id"`
  Name     string `json:"name"`
  Email    string `json:"email"`
  Document string `json:"document"`
  DateOfBirth time.Time `json:"dateOfBirth"`
}

type CreatePatientRequest struct {
  Patient
}

