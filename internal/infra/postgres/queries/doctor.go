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

  UpdateDoctor = `
  UPDATE doctors SET
    name = $1,
    email = $2,
    password = $3,
    document = $4
  WHERE
    id = $5
  `

  DeleteDoctor = `
  DELETE FROM
    doctors
  WHERE
    id = $1
  `
)

