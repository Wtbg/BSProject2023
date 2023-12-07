package device

import (
	"github.com/labstack/echo/v4"
	"go-svc-tpl/app/response"
	"go-svc-tpl/model"
)

func SearchMessageByAttribute(e echo.Context) error {
	var condition MessageSearchCondition
	if err := e.Bind(&condition); err != nil {
		return echo.ErrBadRequest
	}
	if err := e.Validate(&condition); err != nil {
		return e.String(400, err.Error())
	}
	gormCondition := model.DB.Where("device_id = ?", *condition.DeviceID)
	if condition.Alert != nil {
		if *condition.Alert != 0 && *condition.Alert != 1 {
			return e.String(400, "alert must be 0 or 1")
		}
		gormCondition = gormCondition.Where("alert = ?", *condition.Alert)
	}
	if condition.ValueLower != nil {
		if *condition.ValueLower < 0 {
			return e.String(400, "valueLower must be positive")
		}
		gormCondition = gormCondition.Where("value >= ?", *condition.ValueLower)
	}
	if condition.ValueUpper != nil {
		if *condition.ValueUpper < 0 {
			return e.String(400, "valueUpper must be positive")
		}
		gormCondition = gormCondition.Where("value <= ?", *condition.ValueUpper)
	}
	if condition.PageSize != nil {
		if *condition.PageSize < 0 {
			return e.String(400, "PageSize must be positive")
		}
		gormCondition = gormCondition.Limit(*condition.PageSize)
	}
	var messages []model.DeviceMessage
	gormCondition.Find(&messages)
	//fmt.Println(condition)
	return e.JSON(200,
		response.Body{
			Code:   response.OK,
			Msg:    "search message success",
			Result: messages,
		},
	)
}

type MessageSearchCondition struct {
	DeviceID   *string `json:"deviceID" validate:"required"`
	Alert      *int    `json:"alert"`
	ValueLower *int    `json:"valueLower"`
	ValueUpper *int    `json:"valueUpper"`
	PageSize   *int    `json:"pageSize"`
}
