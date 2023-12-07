package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Logout(c echo.Context) error {
	//clean jwt cookie and redirect to login page
	clearCookie := &http.Cookie{
		Name:    "jwt",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
		MaxAge:  -1,
	}
	c.SetCookie(clearCookie)
	return c.Redirect(http.StatusFound, "/login")
}
