package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("JWT_SECRET_KEY")

// Claims defines the structure for JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Generates a new JWT token
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodH256, claims)
	return token.signedString(jwtKey)
}

// Validates the jwt token
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return claims, nil
}
