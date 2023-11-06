package doctor_services

import (
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

