package helpers

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"todo-app/configs"
	"todo-app/models"
)

func GenerateToken(user models.User) string {
	// Create the Claims
	claims := jwt.MapClaims{
		"name":  user.Name,
		"email": user.Email,
		"id":    user.ID,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, _ := token.SignedString([]byte(configs.Env.JwtSecret))
	return t
}
