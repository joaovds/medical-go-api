package entities

type ValidationError struct {
  Message string `json:"message"`
}

type RequestInput interface {
  Validate() ValidationError
}

