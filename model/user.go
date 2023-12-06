package model

type User struct {
	UserID   int    `gorm:"not null;autoIncrement;primaryKey;column:user_id"`
	Username string `gorm:"not null;column:username;unique"`
	Password string `gorm:"not null;column:password"`
	Email    string `gorm:"not null;column:email;unique"`
}
