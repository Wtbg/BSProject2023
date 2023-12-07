package model

type Device struct {
	DeviceID    string `gorm:"not null;primaryKey;column:device_id"`
	DeviceName  string `gorm:"not null;column:device_name;default:default"`
	DeviceType  string `gorm:"not null;column:device_type;default:default"`
	Alert       bool   `gorm:"not null;column:alert;index:alert;default:false"`
	IsConnected bool   `gorm:"not null;column:is_connected;index:is_connected;default:false"`
}
