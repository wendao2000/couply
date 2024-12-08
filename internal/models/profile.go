package model

type BasicProfile struct {
	UserID      int64
	Name        string
	DateOfBirth string
}

type Profile struct {
	ID          int64 `gorm:"primaryKey;autoIncrement;not null;type:INTEGER"`
	UserID      int64 `gorm:"not null"`
	Name        string
	DateOfBirth string `gorm:"type:DATE"`
}
