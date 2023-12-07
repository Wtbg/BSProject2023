package device

import (
	"github.com/labstack/echo/v4"
	"go-svc-tpl/app/response"
	"go-svc-tpl/model"
	"net/http"
)

func SearchDevice(c echo.Context) error {
	var condition DeviceSearchCondition
	if err := c.Bind(&condition); err != nil {
		return echo.ErrBadRequest
	}
	gormCondition := model.DB
	if condition.DeviceName != nil {
		gormCondition = gormCondition.Where("device_name is ?", *condition.DeviceName)
	}
	if condition.DeviceType != nil {
		gormCondition = gormCondition.Where("device_type is ?", *condition.DeviceType)
	}
	if condition.Alert != nil {
		gormCondition = gormCondition.Where("alert is ?", *condition.Alert)
	}
	if condition.IsConnected != nil {
		gormCondition = gormCondition.Where("is_connected is ?", *condition.IsConnected)
	}
	if condition.DeviceID != nil {
		gormCondition = gormCondition.Where("device_id LIKE ?", "%"+*condition.DeviceID+"%")
	}
	var devices []model.Device
	gormCondition.Find(&devices)

	return c.JSON(http.StatusOK,
		response.Body{
			Code:   response.OK,
			Msg:    "search device success",
			Result: devices,
		},
	)

}

type DeviceSearchCondition struct {
	DeviceID    *string `json:"deviceID"`
	DeviceName  *string `json:"deviceName"`
	DeviceType  *string `json:"deviceType"`
	Alert       *bool   `json:"alert"`
	IsConnected *bool   `json:"isConnected"`
}
