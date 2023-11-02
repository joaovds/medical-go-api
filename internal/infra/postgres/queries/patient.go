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

  DeletePatient = `
  DELETE FROM
    patients
  WHERE
    id = $1
  `

  CreatePatient = `
  INSERT INTO patients (
    name,
    email,
    document,
    birth_date
  ) VALUES
    ($1, $2, $3, $4)
  RETURNING
    id,
    name,
    email,
    document,
    birth_date
  `

  UpdatePatient = `
  UPDATE patients SET
    name = $1,
    email = $2,
    document = $3,
    birth_date = $4
  WHERE
    id = $5
  `
)

