package entities

type Doctor struct {
  Id       string `json:"id"`
  Name     string `json:"name"`
  Email    string `json:"email"`
  password string
  Document string `json:"document"`
}

