package model

type User struct {
	ID       int64  `gorm:"primaryKey;autoIncrement;not null;type:INTEGER"`
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password []byte `gorm:"not null"`
}
