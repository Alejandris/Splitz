package service

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"

	"splitz/internal/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MethodService struct{ db *pgxpool.Pool }

func NewMethodService(db *pgxpool.Pool) *MethodService { return &MethodService{db: db} }

func (service *MethodService) List(ctx context.Context, userID uuid.UUID) ([]models.Method, error) {
	rows, err := service.db.Query(ctx, `
		SELECT id, name, is_custom, created_by_user_id
		FROM methods
		WHERE created_by_user_id IS NULL OR created_by_user_id = $1
		ORDER BY is_custom, name`, userID)
	if err != nil {
		return nil, fmt.Errorf("list methods: %w", err)
	}
	defer rows.Close()
	result := make([]models.Method, 0)
	for rows.Next() {
		var item models.Method
		if err := rows.Scan(&item.ID, &item.Name, &item.IsCustom, &item.CreatedByUserID); err != nil {
			return nil, fmt.Errorf("scan method: %w", err)
		}
		result = append(result, item)
	}
	return result, rows.Err()
}

func (service *MethodService) Create(ctx context.Context, userID uuid.UUID, request models.CreateMethodRequest) (models.Method, error) {
	if strings.TrimSpace(request.Name) == "" {
		return models.Method{}, errors.New("name is required")
	}
	if len(request.Allocations) == 0 {
		return models.Method{}, errors.New("allocations are required")
	}
	var total float64
	for _, allocation := range request.Allocations {
		if allocation.CategoryID == uuid.Nil || allocation.Percentage < 0 {
			return models.Method{}, errors.New("invalid allocation")
		}
		total += allocation.Percentage
	}
	if math.Abs(total-100) > 0.0001 {
		return models.Method{}, errors.New("allocations must add up to 100")
	}

	tx, err := service.db.Begin(ctx)
	if err != nil {
		return models.Method{}, fmt.Errorf("begin method: %w", err)
	}
	defer tx.Rollback(ctx)
	var item models.Method
	if err := tx.QueryRow(ctx, `
		INSERT INTO methods (name, is_custom, created_by_user_id)
		VALUES ($1, TRUE, $2)
		RETURNING id, name, is_custom, created_by_user_id`,
		strings.TrimSpace(request.Name), userID).
		Scan(&item.ID, &item.Name, &item.IsCustom, &item.CreatedByUserID); err != nil {
		return models.Method{}, fmt.Errorf("create method: %w", err)
	}
	for _, allocation := range request.Allocations {
		result, err := tx.Exec(ctx, `
			INSERT INTO method_allocations (method_id, category_id, percentage)
			SELECT $1, id, $3 FROM categories
			WHERE id = $2 AND is_active = TRUE
				AND (created_by_user_id IS NULL OR created_by_user_id = $4)`,
			item.ID, allocation.CategoryID, allocation.Percentage, userID)
		if err != nil {
			return models.Method{}, fmt.Errorf("create method allocation: %w", err)
		}
		if result.RowsAffected() != 1 {
			return models.Method{}, errors.New("allocation category not found or not available")
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return models.Method{}, fmt.Errorf("commit method: %w", err)
	}
	return item, nil
}
