package user

import (
	"github.com/labstack/echo/v4"
	"go-svc-tpl/app/response"
	"go-svc-tpl/model"
	"net/http"
)

// @tags Register
// @summary Register
// @router /user/register [post]
// @produce json
// @param username formData string true "username"
// @param email formData string true "email"
// @param password formData string true "password"
// @response 200 {object} response.Body
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
	return c.JSON(http.StatusOK,
		response.Body{
			Code:   response.OK,
			Msg:    "register success",
			Result: user,
		},
	)
}
