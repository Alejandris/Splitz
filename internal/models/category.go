package models

import "github.com/google/uuid"

type Category struct {
	ID              uuid.UUID  `json:"id"`
	Name            string     `json:"name"`
	Type            string     `json:"type"`
	CreatedByUserID *uuid.UUID `json:"created_by_user_id,omitempty"`
	IsActive        bool       `json:"is_active"`
}

type CreateCategoryRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type UpdateCategoryRequest struct {
	Name     string `json:"name"`
	IsActive *bool  `json:"is_active"`
}
