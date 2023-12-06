package model

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	connectDatabase()
	//multiple auto migrate and error handling
	var err error
	err = DB.AutoMigrate(&User{})
	if err != nil {
		logrus.Panic(err)
	}
	err = DB.AutoMigrate(&Device{})
	if err != nil {
		logrus.Panic(err)
	}
	err = DB.AutoMigrate(&DeviceMessage{})
	if err != nil {
		logrus.Panic(err)
	}
}

func connectDatabase() {
	viper.SetConfigName("conf")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic(err)
	}

	loginInfo := viper.GetStringMapString("sql")

	dbArgs := loginInfo["username"] + ":" + loginInfo["password"] +
		"@(localhost)/" + loginInfo["db_name"] + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dbArgs), &gorm.Config{})
	sqlDB, err := DB.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	if err != nil {
		logrus.Panic(err)
	}
}
