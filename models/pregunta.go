package models

import "github.com/google/uuid"

type Pregunta struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;not null;unique;default:gen_random_uuid()" json:"id"`
	Nombre      string    `json:"nombre"`
	Obligatorio bool      `json:"obligatorio"`
}
