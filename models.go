package main

import (
	"time"

	"github.com/LksR-dev/go-course/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKEY    string    `json:"api_key"`
}

func databaseUserToUser(dbUSer database.User) User {
	return User{
		ID:        dbUSer.ID,
		CreatedAt: dbUSer.CreatedAt,
		UpdatedAt: dbUSer.UpdatedAt,
		Name:      dbUSer.Name,
		APIKEY:    dbUSer.ApiKey,
	}
}
