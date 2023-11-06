package entities

type Doctor struct {
  Id       string `json:"id"`
  Name     string `json:"name"`
  Email    string `json:"email"`
  Password string `json:"password"`
  Document string `json:"document"`
}

