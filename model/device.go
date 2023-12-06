package model

type Device struct {
	DeviceID    string `gorm:"not null;primaryKey;column:device_id"`
	DeviceName  string `gorm:"not null;column:device_name"`
	DeviceType  string `gorm:"not null;column:device_type"`
	IsWarning   bool   `gorm:"not null;column:is_warning;index:is_warning"`
	IsConnected bool   `gorm:"not null;column:is_connected;index:is_connected"`
}
