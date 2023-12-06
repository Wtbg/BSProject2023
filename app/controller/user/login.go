package user

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go-svc-tpl/model"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type jwtCustomClaims struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func Login(c echo.Context) error {
	method := c.FormValue("method")
	password := c.FormValue("password")
	var user model.User

	var result *gorm.DB
	if method == "username" {
		username := c.FormValue("username")
		result = model.DB.Where(&model.User{Username: username, Password: password}).First(&user)
	} else if method == "email" {
		email := c.FormValue("email")
		result = model.DB.Where(&model.User{Email: email, Password: password}).First(&user)
	} else {
		return echo.ErrUnauthorized
	}
	// Try finding user in database
	if result.Error != nil {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		user.UserID,
		user.Username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
