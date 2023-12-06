package model

type DeviceMessage struct {
	DeviceID  string  `gorm:"not null;column:device_id;index:device_id"`
	Info      string  `gorm:"not null;column:info"`
	Value     int     `gorm:"not null;column:value"`
	Alert     bool    `gorm:"not null;column:alert"`
	Lng       float64 `gorm:"not null;column:lng"`
	Lat       float64 `gorm:"not null;column:lat"`
	Timestamp int64   `gorm:"not null;column:timestamp;index:timestamp"`
}
