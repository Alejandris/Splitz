package service

import (
	"context"
	"fmt"

	"splitz/internal/auth"
	"splitz/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService struct {
	db *pgxpool.Pool
}

func NewUserService(db *pgxpool.Pool) *UserService {
	return &UserService{db: db}
}

func (service *UserService) Sync(ctx context.Context, identity auth.Identity) (models.User, error) {
	var user models.User
	err := service.db.QueryRow(ctx, `
		INSERT INTO users (firebase_uid, email, name)
		VALUES ($1, $2, NULLIF($3, ''))
		ON CONFLICT (firebase_uid) DO UPDATE SET
			email = EXCLUDED.email,
			name = EXCLUDED.name
		RETURNING id, firebase_uid, email, COALESCE(name, '')`,
		identity.UID, identity.Email, identity.Name,
	).Scan(&user.ID, &user.FirebaseUID, &user.Email, &user.Name)
	if err != nil {
		return models.User{}, fmt.Errorf("sync user: %w", err)
	}
	return user, nil
}
