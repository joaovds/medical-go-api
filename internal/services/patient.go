package patient_services

import (
	"log"

	"github.com/joaovds/first-go-api/internal/entities"
	"github.com/joaovds/first-go-api/internal/infra/postgres"
	"github.com/joaovds/first-go-api/internal/infra/postgres/queries"
)

func GetAllPatients() ([]entities.Patient, error) {
  db, err := postgres.GetConnection()
  if err != nil {
    return nil, err
  }

  patientsRows, err := db.Query(queries.GetAllPatients)
  if err != nil {
    return nil, err
  }
  defer patientsRows.Close()
  defer db.Close()

  var patients []entities.Patient = []entities.Patient{}

  for patientsRows.Next() {
    var patient entities.Patient

    err := patientsRows.Scan(
      &patient.Id,
      &patient.Name,
      &patient.Email,
      &patient.Document,
      &patient.DateOfBirth,
      )
    if err != nil {
      log.Panicln(err)
    }

    patients = append(patients, patient)
  }

  return patients, nil
}

