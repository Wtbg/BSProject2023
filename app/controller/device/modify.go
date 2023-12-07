package device

import (
	"github.com/labstack/echo/v4"
	"go-svc-tpl/app/response"
	"go-svc-tpl/model"
	"net/http"
)

// @tags Modify
// @summary Modify
// @router /device/modify [post]
// @produce json
// @param deviceID formData string true "deviceID"
// @param deviceName formData string true "deviceName"
// @param deviceType formData string true "deviceType"
// @response 200 {object} response.Body
func Modify(c echo.Context) error {
	deviceID := c.FormValue("deviceID")
	deviceName := c.FormValue("deviceName")
	deviceType := c.FormValue("deviceType")

	// Try update device in database
	device := model.Device{
		DeviceID:   deviceID,
		DeviceName: deviceName,
		DeviceType: deviceType,
	}
	//error if not exist
	err := model.DB.Save(&device).Error
	if err != nil {
		return c.JSON(
			http.StatusOK,
			response.Body{
				Code:   response.ErrBadRequest,
				Msg:    "device not exist",
				Result: nil,
			},
		)
	}
	return c.JSON(http.StatusOK,
		response.Body{
			Code:   response.OK,
			Msg:    "modify device success",
			Result: device,
		},
	)

}
