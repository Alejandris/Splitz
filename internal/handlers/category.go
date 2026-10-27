package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"splitz/internal/auth"
	"splitz/internal/models"
	"splitz/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type CategoryHandler struct {
	users      *service.UserService
	categories *service.CategoryService
}

func NewCategoryHandler(users *service.UserService, categories *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{users: users, categories: categories}
}

func (handler *CategoryHandler) currentUser(w http.ResponseWriter, r *http.Request) (models.User, bool) {
	identity, ok := auth.IdentityFromContext(r.Context())
	if !ok {
		writeUnauthorized(w)
		return models.User{}, false
	}
	user, err := handler.users.Sync(r.Context(), identity)
	if err != nil {
		http.Error(w, "could not synchronize user", http.StatusInternalServerError)
		return models.User{}, false
	}
	return user, true
}

func (handler *CategoryHandler) List(w http.ResponseWriter, r *http.Request) {
	user, ok := handler.currentUser(w, r)
	if !ok {
		return
	}
	items, err := handler.categories.List(r.Context(), user.ID)
	if err != nil {
		http.Error(w, "could not load categories", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, items)
}

func (handler *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	user, ok := handler.currentUser(w, r)
	if !ok {
		return
	}
	var request models.CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	item, err := handler.categories.Create(r.Context(), user.ID, request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func (handler *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	user, ok := handler.currentUser(w, r)
	if !ok {
		return
	}
	categoryID, err := uuid.Parse(chi.URLParam(r, "categoryID"))
	if err != nil {
		http.Error(w, "invalid category id", http.StatusBadRequest)
		return
	}
	var request models.UpdateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	item, err := handler.categories.Update(r.Context(), user.ID, categoryID, request)
	if errors.Is(err, service.ErrCategoryNotFound) {
		http.Error(w, "category not found or not editable", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusOK, item)
}
