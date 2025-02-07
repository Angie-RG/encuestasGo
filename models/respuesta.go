package models

import "github.com/google/uuid"

type Respuesta struct {
	ID     uuid.UUID `gorm:"primaryKey;type:uuid;not null;unique;default:gen_random_uuid()" json:"id"`
	Valor  string    `json:"valor"`
	Motivo string    `json:"motivo"`
}
