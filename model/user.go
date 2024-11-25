package model

import "github.com/google/uuid"

type User struct {
	ID              uuid.UUID `json:"id_user"`
	PaternalSurname string    `json:"apellido_paterno"`
	MaternalSurname string    `json:"apellido_materno"`
	Email           string    `json:"email"`
	Names           string    `json:"nombres"`
	Dni             string    `json:"dni"`
	IsAdmin         bool      `json:"is_admin"`
	Password        string    `json:"password"`
	CreatedAt       int64     `json:"created_at"`
	UpdatedAt       int64     `json:"updated_at"`
}

type Users []User
