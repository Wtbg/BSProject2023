package app

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	deviceRouter "go-svc-tpl/app/controller/device"
	userRouter "go-svc-tpl/app/controller/user"
)

type jwtCustomClaims struct {
	Name string `json:"username"`
	ID   int    `json:"user_id"`
	jwt.RegisteredClaims
}

func addRoutes() {
	api := e.Group("api")

	api.GET("/doc/*", echoSwagger.WrapHandler)

	// user group
	user := api.Group("/user")

	user.POST("/register", userRouter.Register)

	user.POST("/login", userRouter.Login)

	user.POST("/logout", userRouter.Logout)

	device := api.Group("/device")

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}

	device.Use(echojwt.WithConfig(config))

	device.POST("/add", deviceRouter.Create)

	device.POST("/modify", deviceRouter.Modify)

	device.POST("/searchMessage", deviceRouter.SearchMessageByAttribute)

	device.POST("/searchDevice", deviceRouter.SearchDevice)

	device.POST("/summary", deviceRouter.DeviceSummary)
}
