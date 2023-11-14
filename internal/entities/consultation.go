package entities

import "time"

type Consultation struct {
  Id          string    `json:"id"`
  Date        time.Time `json:"date"`
  Description string    `json:"description"`
  Notes       string    `json:"notes"`
  Diagnosis   string    `json:"diagnosis"`
  Doctor      Doctor    `json:"doctor"`
  Patient     Patient   `json:"patient"`
}

