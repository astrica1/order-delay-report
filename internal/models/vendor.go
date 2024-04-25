package models

type Vendor struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"null" json:"name"`
	Address     string `gorm:"null" json:"address"`
	PhoneNumber string `gorm:"type:varchar(10);null" json:"phone_number"`
	IsActive    bool   `gorm:"default:true" json:"is_active"`
}
