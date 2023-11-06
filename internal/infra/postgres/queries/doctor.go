package queries

const (
  GetAllDoctors = `
  SELECT
    id,
    name,
    email,
    password,
    document
  FROM
    doctors
  `
)

