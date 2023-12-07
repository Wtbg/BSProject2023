package device

import (
	"github.com/labstack/echo/v4"
	"go-svc-tpl/model"
	"net/http"
)

func DeviceSummary(c echo.Context) error {
	//find device number connected number and is_alerting number
	var deviceCount int64
	var connectedCount int64
	var alertingCount int64
	model.DB.Model(&model.Device{}).Count(&deviceCount)
	model.DB.Model(&model.Device{}).Where("is_connected = ?", true).Count(&connectedCount)
	model.DB.Model(&model.Device{}).Where("alert = ?", true).Count(&alertingCount)
	//find message total number
	var messageCount int64
	model.DB.Model(&model.DeviceMessage{}).Count(&messageCount)
	return c.JSON(http.StatusOK, map[string]int64{
		"deviceCount":    deviceCount,
		"connectedCount": connectedCount,
		"alertingCount":  alertingCount,
		"messageCount":   messageCount,
	},
	)
}
