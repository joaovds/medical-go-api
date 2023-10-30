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
)

