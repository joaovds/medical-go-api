package validation

import "github.com/google/uuid"

func IsValidUUID(id string) bool {
  // verificar se é valido
  _, err := uuid.Parse(id)
  return err == nil
}

