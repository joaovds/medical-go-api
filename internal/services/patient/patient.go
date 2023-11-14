package patient_services

import (
	"database/sql"
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
  defer db.Close() 

  patientsRows, err := db.Query(queries.GetAllPatients)
  if err != nil {
    return nil, err
  }
  defer patientsRows.Close()

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

func GetPatientById(patientId string) (*entities.Patient, error) {
  db, err := postgres.GetConnection()
  if err != nil {
    return nil, err
  }
  defer db.Close() 

  patientRow := db.QueryRow(queries.GetPatientById, patientId)

  patient := new(entities.Patient)

  err = patientRow.Scan(
    &patient.Id,
    &patient.Name,
    &patient.Email,
    &patient.Document,
    &patient.DateOfBirth,
    )
  if err == sql.ErrNoRows {
    return nil, nil
  }
  if err != nil {
    return nil, err
  }

  return patient, nil
}

func CreatePatient(patient *entities.CreatePatientRequest) (*entities.Patient, error) {
  db, err := postgres.GetConnection()
  if err != nil {
    return nil, err
  }
  defer db.Close()

  createdPatient := new(entities.Patient)

  patientRow := db.QueryRow(
    queries.CreatePatient,
    patient.Name,
    patient.Email,
    patient.Document,
    patient.DateOfBirth,
  )

  err = patientRow.Scan(
    &createdPatient.Id,
    &createdPatient.Name,
    &createdPatient.Email,
    &createdPatient.Document,
    &createdPatient.DateOfBirth,
  )
  if err != nil {
    return nil, err
  }

  return createdPatient, nil
}

func UpdatePatient(patientId string, patient *entities.Patient) error {
  db, err := postgres.GetConnection()
  if err != nil {
    return err
  }
  defer db.Close() 

  _, err = db.Exec(
    queries.UpdatePatient,
    patient.Name,
    patient.Email,
    patient.Document,
    patient.DateOfBirth,
    patientId,
  )
  if err != nil {
    return err
  }

  return nil
}

func DeletePatient(patientId string) error {
  db, err := postgres.GetConnection()
  if err != nil {
    return err
  }
  defer db.Close() 

  _, err = db.Exec(queries.DeletePatient, patientId)
  if err != nil {
    return err
  }

  return nil
}

