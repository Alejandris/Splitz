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

type PeriodHandler struct {
	users   *service.UserService
	periods *service.PeriodService
}

func NewPeriodHandler(users *service.UserService, periods *service.PeriodService) *PeriodHandler {
	return &PeriodHandler{users: users, periods: periods}
}

func (handler *PeriodHandler) CreateOrGetPeriod(w http.ResponseWriter, r *http.Request) {
	identity, ok := auth.IdentityFromContext(r.Context())
	if !ok {
		writeUnauthorized(w)
		return
	}
	user, err := handler.users.Sync(r.Context(), identity)
	if err != nil {
		http.Error(w, "could not synchronize user", http.StatusInternalServerError)
		return
	}

	var request models.CreatePeriodRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	period, err := handler.periods.CreateOrGetPeriod(r.Context(), user.ID, request)
	if errors.Is(err, service.ErrPeriodAlreadyExists) {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusOK, period)
}

func (handler *PeriodHandler) ListPeriods(w http.ResponseWriter, r *http.Request) {
	user, ok := handler.currentUser(w, r)
	if !ok {
		return
	}
	periods, err := handler.periods.ListPeriods(r.Context(), user.ID)
	if err != nil {
		http.Error(w, "could not load periods", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, periods)
}

func (handler *PeriodHandler) GetPeriod(w http.ResponseWriter, r *http.Request) {
	user, ok := handler.currentUser(w, r)
	if !ok {
		return
	}
	periodID, err := uuid.Parse(chi.URLParam(r, "periodID"))
	if err != nil {
		http.Error(w, "invalid period id", http.StatusBadRequest)
		return
	}
	period, err := handler.periods.GetPeriod(r.Context(), user.ID, periodID)
	if errors.Is(err, service.ErrPeriodNotFound) {
		http.Error(w, "period not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "could not load period", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, period)
}

func (handler *PeriodHandler) ClosePeriod(w http.ResponseWriter, r *http.Request) {
	user, ok := handler.currentUser(w, r)
	if !ok {
		return
	}
	periodID, err := uuid.Parse(chi.URLParam(r, "periodID"))
	if err != nil {
		http.Error(w, "invalid period id", http.StatusBadRequest)
		return
	}
	period, err := handler.periods.ClosePeriod(r.Context(), user.ID, periodID)
	if errors.Is(err, service.ErrPeriodNotFound) {
		http.Error(w, "period not found", http.StatusNotFound)
		return
	}
	if errors.Is(err, service.ErrPeriodClosed) {
		http.Error(w, "period is already closed", http.StatusConflict)
		return
	}
	if err != nil {
		http.Error(w, "could not close period", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, period)
}

func (handler *PeriodHandler) CreateFixedExpense(w http.ResponseWriter, r *http.Request) {
	user, ok := handler.currentUser(w, r)
	if !ok {
		return
	}
	var request models.CreateFixedExpenseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	expense, err := handler.periods.CreateFixedExpense(r.Context(), user.ID, request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusCreated, expense)
}

func (handler *PeriodHandler) ListFixedExpenses(w http.ResponseWriter, r *http.Request) {
	user, ok := handler.currentUser(w, r)
	if !ok {
		return
	}
	expenses, err := handler.periods.ListFixedExpenses(r.Context(), user.ID)
	if err != nil {
		http.Error(w, "could not load fixed expenses", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, expenses)
}

func (handler *PeriodHandler) ListHistoricalFixedExpenses(w http.ResponseWriter, r *http.Request) {
	user, ok := handler.currentUser(w, r)
	if !ok {
		return
	}
	expenses, err := handler.periods.ListHistoricalFixedExpenses(r.Context(), user.ID)
	if err != nil {
		http.Error(w, "could not load historical fixed expenses", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, expenses)
}

func (handler *PeriodHandler) UpdateFixedExpense(w http.ResponseWriter, r *http.Request) {
	user, ok := handler.currentUser(w, r)
	if !ok {
		return
	}
	expenseID, err := uuid.Parse(chi.URLParam(r, "expenseID"))
	if err != nil {
		http.Error(w, "invalid expense id", http.StatusBadRequest)
		return
	}
	var request models.UpdateFixedExpenseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	expense, err := handler.periods.UpdateFixedExpense(r.Context(), user.ID, expenseID, request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusOK, expense)
}

func (handler *PeriodHandler) SaveFinancialSummary(w http.ResponseWriter, r *http.Request) {
	user, ok := handler.currentUser(w, r)
	if !ok {
		return
	}
	periodID, err := uuid.Parse(chi.URLParam(r, "periodID"))
	if err != nil {
		http.Error(w, "invalid period id", http.StatusBadRequest)
		return
	}
	var request models.CreateFinancialSummaryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	summary, err := handler.periods.UpsertFinancialSummary(r.Context(), user.ID, periodID, request)
	if errors.Is(err, service.ErrPeriodNotFound) {
		http.Error(w, "period not found", http.StatusNotFound)
		return
	}
	if errors.Is(err, service.ErrPeriodClosed) {
		http.Error(w, "period is closed", http.StatusConflict)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusOK, summary)
}

func (handler *PeriodHandler) PreviewFinancialSummary(w http.ResponseWriter, r *http.Request) {
	user, ok := handler.currentUser(w, r)
	if !ok {
		return
	}
	periodID, err := uuid.Parse(chi.URLParam(r, "periodID"))
	if err != nil {
		http.Error(w, "invalid period id", http.StatusBadRequest)
		return
	}
	var request models.CreateFinancialSummaryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	summary, err := handler.periods.PreviewFinancialSummary(r.Context(), user.ID, periodID, request)
	if errors.Is(err, service.ErrPeriodNotFound) {
		http.Error(w, "period not found", http.StatusNotFound)
		return
	}
	if errors.Is(err, service.ErrPeriodClosed) {
		http.Error(w, "period is closed", http.StatusConflict)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusOK, summary)
}

func (handler *PeriodHandler) currentUser(w http.ResponseWriter, r *http.Request) (models.User, bool) {
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

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func writeUnauthorized(w http.ResponseWriter) {
	writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
}
