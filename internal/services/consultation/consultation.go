package consultation_services

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

func CreateConsultation(consultation *entities.CreateConsultationRequest) error {
  db, err := postgres.GetConnection()
  if err != nil {
    return err
  }
  defer db.Close()

  fmt.Println(consultation)

  createdConsultation := db.QueryRow(
    queries.CreateConsultation,
    consultation.Date,
    consultation.Description,
    consultation.Notes,
    consultation.Diagnosis,
    consultation.DoctorId,
    consultation.PatientId,
    )

  err = createdConsultation.Err();
  if err != nil {
    log.Println(err)
    return fmt.Errorf("Error creating consultation")
  }

  return nil
}

func UpdateConsultation(consultationId string, consultation *entities.UpdateConsultationRequest) error {
  db, err := postgres.GetConnection()
  if err != nil {
    return err
  }
  defer db.Close() 

  _, err = db.Exec(
    queries.UpdateConsultation,
    consultation.Date,
    consultation.Description,
    consultation.Notes,
    consultation.Diagnosis,
    consultation.DoctorId,
    consultation.PatientId,
    consultationId,
    )
  if err != nil {
    return err
  }

  return nil
}

func DeleteConsultation(consultationId string) error {
  db, err := postgres.GetConnection()
  if err != nil {
    return err
  }
  defer db.Close() 

  _, err = db.Exec(queries.DeleteConsultation, consultationId)
  if err != nil {
    return err
  }

  return nil
}

