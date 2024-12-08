package model

// TODO: ID use UUID
type User struct {
	ID       int64  `gorm:"primaryKey;autoIncrement;not null;type:INTEGER"`
	Username string `gorm:"uniqueIndex:idx_users_username,where:username != ''"`
	Email    string `gorm:"uniqueIndex:idx_users_email,where:email != ''"`
	Password []byte `gorm:"not null"`
}
