package auth

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/namburisnehitha/IssueTracker/domain"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(Password string) (string, error) {

	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(Password), 14)

	if err != nil {
		return "", err
	}

	return string(hashedpassword), err
}

func CheckPasswordHash(password, hashedpassword string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
	return err == nil
}

func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}
	return []byte(secret)
}

func GenerateToken(userID string) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getJWTSecret())

}

func ReturnAsToken(token *jwt.Token) (interface{}, error) {
	return getJWTSecret(), nil
}

func ValidateToken(tokenString string) (string, error) {

	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, ReturnAsToken)

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*jwt.MapClaims)

	if !ok || !token.Valid {
		return "", domain.ErrInvalidToken
	}

	userID := (*claims)["user_id"].(string)

	return userID, err
}

func JWTMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		authToken := r.Header.Get("Authorization")

		if authToken == "" {
			http.Error(w, domain.ErrEmptyHeader.Error(), http.StatusBadRequest)
			return
		}

		tokenString := strings.TrimPrefix(authToken, "Bearer ")
		userID, err := ValidateToken(tokenString)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), domain.UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))

	}
	return http.HandlerFunc(fn)
}
