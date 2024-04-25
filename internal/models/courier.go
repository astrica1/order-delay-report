package models

import "time"

type Courier struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	FullName     string    `gorm:"null" json:"full_name"`
	PhoneNumber  string    `gorm:"type:varchar(10);null" json:"phone_number"`
	NationalID   string    `gorm:"type:varchar(10);null" json:"national_id"`
	BirthDate    time.Time `gorm:"null" json:"birth_date"`
	LicensePlate string    `gorm:"null" json:"license_plate"`
	IsActive     bool      `gorm:"default:true" json:"is_active"`
}
