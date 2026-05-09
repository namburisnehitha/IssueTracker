package auth

import (
	"net/http"
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

func GenerateToken(userID string) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("9q2CLgDY"))

}

func ReturnAsToken(token *jwt.Token) (interface{}, error) {
	return []byte("9q2CLgDY"), nil
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
		_, err := ValidateToken(tokenString)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)

	}
	return http.HandlerFunc(fn)
}
