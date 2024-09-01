package models

import (
	"time"
)

// User represents a user in the system
// @Description User model with GORM fields
// @Accept json
// @Produce json
type User struct {
	ID           uint      `json:"id"`           // @Description The unique identifier for the user
	CreatedAt    time.Time `json:"createdAt"`    // @Description Timestamp when the user was created
	UpdatedAt    time.Time `json:"updatedAt"`    // @Description Timestamp when the user was last updated
	DeletedAt    time.Time `json:"deletedAt"`    // @Description Timestamp when the user was deleted
	UsuarioNome  string    `json:"usuarioNome"`  // @Description User's full name
	UsuarioLogin string    `json:"usuarioLogin"` // @Description User's login
	Senha        string    `json:"senha"`        // @Description User's password
}
