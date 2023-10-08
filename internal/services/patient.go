package patient_services

import (
	"time"

	"github.com/joaovds/first-go-api/internal/entities"
)

func GetAllPatients() ([]entities.Patient, error) {
  patients := []entities.Patient{
    {
      Id: "1234",
      Name: "John Doe",
      Email: "john@email.com",
      Document: "123456789",
      DateOfBirth: time.Now(),
    },
  }

  return patients, nil
}

