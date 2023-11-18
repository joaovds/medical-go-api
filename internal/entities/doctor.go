package entities

import "net/mail"

type Doctor struct {
  Id       string `json:"id"`
  Name     string `json:"name"`
  Email    string `json:"email"`
  Password string `json:"password"`
  Document string `json:"document"`
}

type CreateDoctorRequest struct {
  Doctor
}


func (c *CreateDoctorRequest) Validate() ValidationError {
  if c.Name == "" {
    return ValidationError{Message: "Name is required"}
  }

  if c.Email == "" {
    return ValidationError{Message: "Email is required"}
  }

  _, err := mail.ParseAddress(c.Email)
  if err != nil {
    return ValidationError{Message: "Email is invalid"}
  }
  
  if c.Password == "" {
    return ValidationError{Message: "Password is required"}
  }

  if c.Document == "" {
    return ValidationError{Message: "Document is required"}
  }

  return ValidationError{}
}
