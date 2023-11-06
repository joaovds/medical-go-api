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

  GetDoctorById = GetAllDoctors + `
  WHERE
  id = $1
  `
)

