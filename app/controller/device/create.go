package device

import (
	"github.com/labstack/echo/v4"
	"go-svc-tpl/app/response"
	"go-svc-tpl/model"
	"net/http"
)

// @tags Create
// @summary Create
// @router /device/add [post]
// @produce json
// @param deviceID formData string true "deviceID"
// @param deviceName formData string false "deviceName"
// @param deviceType formData string false "deviceType"
// @response 200 {object} response.Body
func Create(c echo.Context) error {
	deviceID := c.FormValue("deviceID")
	deviceName := c.FormValue("deviceName")
	deviceType := c.FormValue("deviceType")

	// Try insert device in database
	device := model.Device{
		DeviceID:   deviceID,
		DeviceName: deviceName,
		DeviceType: deviceType,
	}
	//error if already exist
	err := model.DB.Create(&device).Error
	if err != nil {
		return c.JSON(
			http.StatusOK,
			response.Body{
				Code:   response.ErrBadRequest,
				Msg:    "device already exist",
				Result: nil,
			},
		)
	}
	return c.JSON(http.StatusOK,
		response.Body{
			Code:   response.OK,
			Msg:    "create device success",
			Result: device,
		},
	)
}
