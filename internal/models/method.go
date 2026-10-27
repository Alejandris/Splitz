package models

import "github.com/google/uuid"

type Method struct {
	ID              uuid.UUID  `json:"id"`
	Name            string     `json:"name"`
	IsCustom        bool       `json:"is_custom"`
	CreatedByUserID *uuid.UUID `json:"created_by_user_id,omitempty"`
}

type MethodAllocation struct {
	CategoryID uuid.UUID `json:"category_id"`
	Percentage float64   `json:"percentage"`
}

type CreateMethodRequest struct {
	Name        string             `json:"name"`
	Allocations []MethodAllocation `json:"allocations"`
}
