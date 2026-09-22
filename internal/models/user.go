package models

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID `json:"id"`
	FirebaseUID string    `json:"firebase_uid"`
	Email       string    `json:"email"`
	Name        string    `json:"name,omitempty"`
}
