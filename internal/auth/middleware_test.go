package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	firebaseauth "firebase.google.com/go/v4/auth"
)

type tokenVerifierFunc func(context.Context, string) (*firebaseauth.Token, error)

func (verifier tokenVerifierFunc) VerifyIDToken(ctx context.Context, token string) (*firebaseauth.Token, error) {
	return verifier(ctx, token)
}

func TestRequireFirebaseUserRejectsMissingBearerToken(t *testing.T) {
	handler := RequireFirebaseUser(tokenVerifierFunc(func(context.Context, string) (*firebaseauth.Token, error) {
		t.Fatal("no debe validar un token ausente")
		return nil, nil
	}))(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Fatal("no debe ejecutar el handler protegido")
	}))

	recording := httptest.NewRecorder()
	handler.ServeHTTP(recording, httptest.NewRequest(http.MethodGet, "/protected", nil))

	if recording.Code != http.StatusUnauthorized {
		t.Fatalf("se esperaba 401, se obtuvo %d", recording.Code)
	}
}

func TestRequireFirebaseUserRejectsInvalidToken(t *testing.T) {
	handler := RequireFirebaseUser(tokenVerifierFunc(func(context.Context, string) (*firebaseauth.Token, error) {
		return nil, context.DeadlineExceeded
	}))(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Fatal("no debe ejecutar el handler protegido")
	}))

	request := httptest.NewRequest(http.MethodGet, "/protected", nil)
	request.Header.Set("Authorization", "Bearer invalid-token")
	recording := httptest.NewRecorder()
	handler.ServeHTTP(recording, request)

	if recording.Code != http.StatusUnauthorized {
		t.Fatalf("se esperaba 401, se obtuvo %d", recording.Code)
	}
}

func TestRequireFirebaseUserAddsValidatedIdentityToContext(t *testing.T) {
	handler := RequireFirebaseUser(tokenVerifierFunc(func(_ context.Context, token string) (*firebaseauth.Token, error) {
		if token != "valid-token" {
			t.Fatalf("token inesperado: %s", token)
		}
		return &firebaseauth.Token{UID: "firebase-user-1", Claims: map[string]interface{}{
			"email": "user@example.com",
			"name":  "Usuario Splitz",
		}}, nil
	}))(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		identity, ok := IdentityFromContext(request.Context())
		if !ok {
			t.Fatal("no se encontró identidad validada en el contexto")
		}
		if identity.UID != "firebase-user-1" || identity.Email != "user@example.com" || identity.Name != "Usuario Splitz" {
			t.Fatalf("identidad incorrecta: %+v", identity)
		}
		w.WriteHeader(http.StatusNoContent)
	}))

	request := httptest.NewRequest(http.MethodGet, "/protected", nil)
	request.Header.Set("Authorization", "Bearer valid-token")
	recording := httptest.NewRecorder()
	handler.ServeHTTP(recording, request)

	if recording.Code != http.StatusNoContent {
		t.Fatalf("se esperaba 204, se obtuvo %d", recording.Code)
	}
}
