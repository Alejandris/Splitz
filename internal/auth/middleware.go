package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	firebaseauth "firebase.google.com/go/v4/auth"
)

type TokenVerifier interface {
	VerifyIDToken(context.Context, string) (*firebaseauth.Token, error)
}

type Identity struct {
	UID   string
	Email string
	Name  string
}

type identityContextKey struct{}

func RequireFirebaseUser(client TokenVerifier) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			value := strings.TrimSpace(r.Header.Get("Authorization"))
			parts := strings.SplitN(value, " ", 2)
			if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") || strings.TrimSpace(parts[1]) == "" {
				writeUnauthorized(w)
				return
			}

			token, err := client.VerifyIDToken(r.Context(), strings.TrimSpace(parts[1]))
			if err != nil || token.UID == "" {
				writeUnauthorized(w)
				return
			}

			identity := Identity{UID: token.UID, Email: claimString(token.Claims, "email"), Name: claimString(token.Claims, "name")}
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), identityContextKey{}, identity)))
		})
	}
}

func IdentityFromContext(ctx context.Context) (Identity, bool) {
	identity, ok := ctx.Value(identityContextKey{}).(Identity)
	return identity, ok
}

func claimString(claims map[string]interface{}, key string) string {
	value, _ := claims[key].(string)
	return value
}

func writeUnauthorized(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
}
