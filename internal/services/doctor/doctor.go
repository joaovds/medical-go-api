package doctor_services

import (
	"database/sql"
	"log"

	"github.com/joaovds/first-go-api/internal/entities"
	"github.com/joaovds/first-go-api/internal/infra/postgres"
	"github.com/joaovds/first-go-api/internal/infra/postgres/queries"
)

func GetAllDoctors() ([]entities.Doctor, error) {
  db, err := postgres.GetConnection()
  if err != nil {
    return nil, err
  }
  defer db.Close() 

  doctorsRows, err := db.Query(queries.GetAllDoctors)
  if err != nil {
    return nil, err
  }
  defer doctorsRows.Close()

  var doctors []entities.Doctor = []entities.Doctor{}

  for doctorsRows.Next() {
    var doctor entities.Doctor

    err := doctorsRows.Scan(
      &doctor.Id,
      &doctor.Name,
      &doctor.Email,
      &doctor.Document,
      &doctor.Password,
      )
    if err != nil {
      log.Panicln(err)
    }

    doctors = append(doctors, doctor)
  }

  return doctors, nil
}

func GetDoctorById(doctorId string) (*entities.Doctor, error) {
  db, err := postgres.GetConnection()
  if err != nil {
    return nil, err
  }
  defer db.Close() 

  doctorRow := db.QueryRow(queries.GetDoctorById, doctorId)

  doctor := new(entities.Doctor)

  err = doctorRow.Scan(
    &doctor.Id,
    &doctor.Name,
    &doctor.Email,
    &doctor.Password,
    &doctor.Document,
    )
  if err == sql.ErrNoRows {
    return nil, nil
  }
  if err != nil {
    return nil, err
  }

  return doctor, nil
}

func CreateDoctor(doctor *entities.CreateDoctorRequest) (*entities.Doctor, error) {
  db, err := postgres.GetConnection()
  if err != nil {
    return nil, err
  }
  defer db.Close()

  createdDoctor := new(entities.Doctor)

  doctorRow := db.QueryRow(
    queries.CreateDoctor,
    doctor.Name,
    doctor.Email,
    doctor.Password,
    doctor.Document,
  )

  err = doctorRow.Scan(
    &createdDoctor.Id,
    &createdDoctor.Name,
    &createdDoctor.Email,
    &createdDoctor.Password,
    &createdDoctor.Document,
  )
  if err != nil {
    return nil, err
  }

  return createdDoctor, nil
}

func UpdateDoctor(doctorId string, doctor *entities.Doctor) error {
  db, err := postgres.GetConnection()
  if err != nil {
    return err
  }
  defer db.Close() 

  _, err = db.Exec(
    queries.UpdateDoctor,
    doctor.Name,
    doctor.Email,
    doctor.Password,
    doctor.Document,
    doctorId,
  )
  if err != nil {
    return err
  }

  return nil
}

func DeleteDoctor(doctorId string) error {
  db, err := postgres.GetConnection()
  if err != nil {
    return err
  }
  defer db.Close() 

  _, err = db.Exec(queries.DeleteDoctor, doctorId)
  if err != nil {
    return err
  }

  return nil
}

