package models

type Agent struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string `json:"username"`
	Password string `gorm:"type:varchar(32)" json:"password"` //store passwords as MD5 hash
	IsActive bool   `gorm:"default:true" json:"is_active"`
}
