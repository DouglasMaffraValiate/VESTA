package user

import (
	"time"
	"uuid"

	"github.com/DouglasMaffraValiate/VESTA/internal/email"
	"github.com/DouglasMaffraValiate/VESTA/internal/password"
)

type User struct {
	ID                   uuid.UUID
	Email                email.Email
	IsVerified, IsAtive  bool
	Password             password.Password
	CreatedAt, UpdatedAt time.Time
}
