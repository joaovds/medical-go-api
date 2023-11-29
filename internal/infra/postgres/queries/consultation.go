package queries

const (
  GetAllConsultations = `
  SELECT
      con."id",
      con."date",
      con."description",
      con."notes",
      con."diagnosis",
      jsonb_build_object(
          'id',
          doc."id",
          'name',
          doc."name",
          'email',
          doc."email",
          'document',
          doc."document"
      ) AS doctor,
      jsonb_build_object(
          'id',
          pat."id",
          'name',
          pat."name",
          'email',
          pat."email",
          'document',
          pat."document",
          'dateOfBirth',
          pat."birth_date"
      ) AS patient
  FROM
      consultations con
  INNER JOIN
  doctors doc ON
      con."doctor_id" = doc."id"
  INNER JOIN
  patients pat ON
      con."patient_id" = pat."id"
  `

  GetConsultationById = GetAllConsultations + `
  WHERE
    con."id" = $1
  `

  CreateConsultation = `
  INSERT INTO consultations (
    date,
    description,
    notes,
    diagnosis,
    doctor_id,
    patient_id
  ) VALUES
    ($1, $2, $3, $4, $5, $6)
  `

  UpdateConsultation = `
  UPDATE consultations SET
    date = $1,
    description = $2,
    notes = $3,
    diagnosis = $4,
    doctor_id = $5,
    patient_id = $6
  WHERE
    id = $7
  `

  DeleteConsultation = `
  DELETE FROM
    consultations
  WHERE
    id = $1
  `
)

