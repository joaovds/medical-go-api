package entities

type Patient struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Document    string `json:"document"`
	DateOfBirth string `json:"dateOfBirth"`
}

type CreatePatientRequest struct {
	Patient
}
