package models

import "time"

type Order struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID   int       `gorm:"null" json:"customer_id"`
	VendorID     int       `gorm:"null" json:"vendor_id"`
	Items        string    `gorm:"null" json:"items"` //not implemented :)
	DeliveryTime time.Time `gorm:"not null" json:"delivery_time"`

	Customer Customer `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Vendor   Vendor   `gorm:"foreignKey:VendorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
