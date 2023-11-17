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
      con."doctorId" = doc."id"
  INNER JOIN
  patients pat ON
      con."patientId" = pat."id"
  `

  GetConsultationById = GetAllConsultations + `
  WHERE
    con."id" = $1
  `
)

