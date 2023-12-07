package model

type DeviceMessage struct {
	MessageID uint    `gorm:"primaryKey;column:message_id;autoIncrement" json:"messageId"`
	DeviceID  string  `gorm:"not null;column:device_id;index:device_id" json:"clientId"`
	Info      string  `gorm:"not null;column:info" json:"info"`
	Value     int     `gorm:"not null;column:value" json:"value"`
	Alert     int     `gorm:"not null;column:alert" json:"alert"`
	Lng       float64 `gorm:"not null;column:lng" json:"lng"`
	Lat       float64 `gorm:"not null;column:lat" json:"lat"`
	Timestamp int64   `gorm:"not null;column:timestamp;index:timestamp" json:"timestamp"`
}
