package consultation_services

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/joaovds/first-go-api/internal/entities"
	"github.com/joaovds/first-go-api/internal/infra/postgres"
	"github.com/joaovds/first-go-api/internal/infra/postgres/queries"
)

func GetAllConsultations() ([]entities.Consultation, error) {
  db, err := postgres.GetConnection()
  if err != nil {
    return nil, err
  }
  defer db.Close() 

  consultationsRows, err := db.Query(queries.GetAllConsultations)
  if err != nil {
    return nil, err
  }
  defer consultationsRows.Close()

  var consultations []entities.Consultation = []entities.Consultation{}

  for consultationsRows.Next() {
    consultation := new(entities.Consultation)
    var doctorJSON string
    var patientJSON string

    err := consultationsRows.Scan(
      &consultation.Id,
      &consultation.Date,
      &consultation.Description,
      &consultation.Notes,
      &consultation.Diagnosis,
      &doctorJSON,
      &patientJSON,
    )
    if err != nil {
      log.Panicln(err)
    }

    err = json.Unmarshal([]byte(doctorJSON), &consultation.Doctor)
    if err != nil {
      log.Panicln(err)
    }

    err = json.Unmarshal([]byte(patientJSON), &consultation.Patient)
    if err != nil {
      log.Panicln(err)
    }

    consultations = append(consultations, *consultation)
  }

  return consultations, nil
}

func GetConsultationById(consultationId string) (*entities.Consultation, error) {
  db, err := postgres.GetConnection()
  if err != nil {
    return nil, err
  }
  defer db.Close() 

  consultationRow := db.QueryRow(queries.GetConsultationById, consultationId)

  consultation := new(entities.Consultation)
  var doctorJSON string
  var patientJSON string

  err = consultationRow.Scan(
    &consultation.Id,
    &consultation.Date,
    &consultation.Description,
    &consultation.Notes,
    &consultation.Diagnosis,
    &doctorJSON,
    &patientJSON,
  )
  if err == sql.ErrNoRows {
    return nil, nil
  }
  if err != nil {
    return nil, err
  }

  err = json.Unmarshal([]byte(doctorJSON), &consultation.Doctor)
  if err != nil {
    log.Panicln(err)
  }

  err = json.Unmarshal([]byte(patientJSON), &consultation.Patient)
  if err != nil {
    log.Panicln(err)
  }

  return consultation, nil
}

