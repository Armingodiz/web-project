package utils

import (
	"errors"
	"time"

	"web-project/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GetUserName returns the username from the context
func GetUserName(c *gin.Context) (string, error) {
	userName, ok := c.Get("user_name")
	if !ok {
		return "", errors.New("claim not found")
	}
	return userName.(string), nil
}

// ValidateAndGetClaims validates a token and returns the claims
func ValidateAndGetClaims(tokenString string) (jwt.MapClaims, error) {
	secret := []byte(config.Configs.App.JwtSecret)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signature")
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}

}

// CreateToken creates a new token for a specific username and duration
func CreateJwtToken(userName string, duration time.Duration) (string, error) {
	type Claims struct {
		UserName string `json:"user_name"`
		jwt.StandardClaims
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	})
	return jwtToken.SignedString([]byte(config.Configs.App.JwtSecret))
}
