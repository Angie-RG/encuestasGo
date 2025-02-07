package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;not null;unique;defeault:gen_random_uuid()" json:"id"`
	FirstName *string   `gorm:"size:100" json:"first_name"`
	LastName  *string   `gorm:"size:100" json:"last_name"`
	Email     string    `gorm:"size:100;not null" json:"email"`
	Password  string    `gorm:"size:255, not null" json:"password"`
	Active    *bool     `gorm:"not null;default:false" json:"active"`
	CreatedAt time.Time `gorm:"not null;default:clock_timestamp()" json:"-"`
	UpdatedAt time.Time `gorm:"not null;default:clock_timestamp()" json:"-"`
	DeletedAt time.Time `gorm:"index" json:"deleted_at"`
}
