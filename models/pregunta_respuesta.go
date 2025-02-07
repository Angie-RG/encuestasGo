package models

import "github.com/google/uuid"

type PreguntaRespuesta struct {
	PreguntaID  uuid.UUID `gorm:"not null" json:"pregunta_id"`
	Pregunta    Pregunta  `json:"pregunta"`
	RespuestaID uuid.UUID `gorm:"not null" json:"respuesta_id"`
	Respuesta   Respuesta `json:"respuesta"`
}
