package user

import (
	"github.com/labstack/echo/v4"
	"go-svc-tpl/model"
)

func Register(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	email := c.FormValue("email")

	// Try insert user in database
	user := model.User{
		Username: username,
		Password: password,
		Email:    email,
	}
	result := model.DB.Create(&user)
	if result.Error != nil {
		return echo.ErrBadRequest
	}
	return c.String(200, "ok")
}
