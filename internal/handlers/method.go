package handlers

import (
	"encoding/json"
	"net/http"

	"splitz/internal/auth"
	"splitz/internal/models"
	"splitz/internal/service"
)

type MethodHandler struct {
	users   *service.UserService
	methods *service.MethodService
}

func NewMethodHandler(users *service.UserService, methods *service.MethodService) *MethodHandler {
	return &MethodHandler{users: users, methods: methods}
}

func (handler *MethodHandler) currentUser(w http.ResponseWriter, r *http.Request) (models.User, bool) {
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

func (handler *MethodHandler) List(w http.ResponseWriter, r *http.Request) {
	user, ok := handler.currentUser(w, r)
	if !ok {
		return
	}
	items, err := handler.methods.List(r.Context(), user.ID)
	if err != nil {
		http.Error(w, "could not load methods", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, items)
}

func (handler *MethodHandler) Create(w http.ResponseWriter, r *http.Request) {
	user, ok := handler.currentUser(w, r)
	if !ok {
		return
	}
	var request models.CreateMethodRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	item, err := handler.methods.Create(r.Context(), user.ID, request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusCreated, item)
}
