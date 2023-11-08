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

  CreateDoctor = `
  INSERT INTO doctors (
    name,
    email,
    password,
    document
  ) VALUES
    ($1, $2, $3, $4)
  RETURNING
    id,
    name,
    email,
    password,
    document
  `
)

