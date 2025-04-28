package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/zGuiOs/poupeme-server/src/config"
)

// CreateToken create an token with permissions
func CreateToken(UUID string) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["expiration"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["uuid"] = UUID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}

// ValidateToken verify if the token is valid
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)

	token, err := jwt.Parse(tokenString, returnKeyVerification)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token inválido")
}

// ExtractUUID return the uuid from the token in req
func ExtractUUID(r *http.Request) (string, error) {
	tokenString := extractToken(r)

	token, err := jwt.Parse(tokenString, returnKeyVerification)
	if err != nil {
		return "", err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uuid, ok := permissions["uuid"].(string)
		if !ok {
			return "", errors.New("UUID não encontrado no token")
		}
		return uuid, nil
	}

	return "", errors.New("Token inválido")
}

// extractToken from the header of req
func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

// returnKeyVerification verify if the token use the same sign method and then return the secret key
func returnKeyVerification(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
