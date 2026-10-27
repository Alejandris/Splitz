package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"splitz/internal/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrCategoryNotFound = errors.New("category not found")
var ErrGlobalCategory = errors.New("global categories cannot be edited")

type CategoryService struct{ db *pgxpool.Pool }

func NewCategoryService(db *pgxpool.Pool) *CategoryService { return &CategoryService{db: db} }

func (service *CategoryService) List(ctx context.Context, userID uuid.UUID) ([]models.Category, error) {
	rows, err := service.db.Query(ctx, `
		SELECT id, name, type, created_by_user_id, is_active
		FROM categories
		WHERE is_active = TRUE AND (created_by_user_id IS NULL OR created_by_user_id = $1)
		ORDER BY created_by_user_id NULLS FIRST, name`, userID)
	if err != nil {
		return nil, fmt.Errorf("list categories: %w", err)
	}
	defer rows.Close()
	result := make([]models.Category, 0)
	for rows.Next() {
		var item models.Category
		if err := rows.Scan(&item.ID, &item.Name, &item.Type, &item.CreatedByUserID, &item.IsActive); err != nil {
			return nil, fmt.Errorf("scan category: %w", err)
		}
		result = append(result, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate categories: %w", err)
	}
	return result, nil
}

func (service *CategoryService) Create(ctx context.Context, userID uuid.UUID, request models.CreateCategoryRequest) (models.Category, error) {
	name := strings.TrimSpace(request.Name)
	if name == "" {
		return models.Category{}, errors.New("name is required")
	}
	if request.Type != "INCOME" && request.Type != "EXPENSE" && request.Type != "SAVING" && request.Type != "DEBT" {
		return models.Category{}, errors.New("type must be INCOME, EXPENSE, SAVING, or DEBT")
	}
	var item models.Category
	err := service.db.QueryRow(ctx, `
		INSERT INTO categories (name, type, created_by_user_id, is_active)
		VALUES ($1, $2, $3, TRUE)
		RETURNING id, name, type, created_by_user_id, is_active`, name, request.Type, userID).
		Scan(&item.ID, &item.Name, &item.Type, &item.CreatedByUserID, &item.IsActive)
	if err != nil {
		return models.Category{}, fmt.Errorf("create category: %w", err)
	}
	return item, nil
}

func (service *CategoryService) Update(ctx context.Context, userID, categoryID uuid.UUID, request models.UpdateCategoryRequest) (models.Category, error) {
	name := strings.TrimSpace(request.Name)
	if name == "" {
		return models.Category{}, errors.New("name is required")
	}
	var item models.Category
	err := service.db.QueryRow(ctx, `
		UPDATE categories
		SET name = $1, is_active = COALESCE($2, is_active)
		WHERE id = $3 AND created_by_user_id = $4
		RETURNING id, name, type, created_by_user_id, is_active`,
		name, request.IsActive, categoryID, userID).
		Scan(&item.ID, &item.Name, &item.Type, &item.CreatedByUserID, &item.IsActive)
	if errors.Is(err, pgx.ErrNoRows) {
		return models.Category{}, ErrCategoryNotFound
	}
	if err != nil {
		return models.Category{}, fmt.Errorf("update category: %w", err)
	}
	return item, nil
}
