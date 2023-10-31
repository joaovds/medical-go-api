package queries

const (
  GetAllPatients = `
  SELECT
    id,
    name,
    email,
    document,
    birth_date
  FROM
    patients
  `

  GetPatientById = GetAllPatients + `
  WHERE
    id = $1
  `
)

