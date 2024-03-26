package helper

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(email string, userID uint) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"sub":   userID,
		"exp":   time.Now().Add(1 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("rahasia"))
}

func GetJWTClaims(tokenString string) (map[string]any, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid method")
		}
		return []byte("rahasia"), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil

}

func GetSubClaims(claims interface{}) (interface{}, error) {

	mapClaims, ok := claims.(map[string]interface{})

	if !ok {
		return nil, errors.New("not map")
	}

	sub, ok := mapClaims["sub"]
	if !ok {
		return nil, errors.New("not found")
	}

	return sub, nil
}
