package models

type Customer struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	FullName    string `gorm:"null" json:"full_name"`
	PhoneNumber string `gorm:"type:varchar(10);null" json:"phone_number"`
}
