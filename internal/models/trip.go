package models

type TripStatus uint8

const (
	TripStatusAssigned TripStatus = iota
	TripStatusAtVendor
	TripStatusPicked
	TripStatusDelivered
)

type Trip struct {
	ID        int        `gorm:"primaryKey;autoIncrement" json:"id"`
	CourierID int        `gorm:"null" json:"courier_id"`
	OrderID   int        `gorm:"null" json:"order_id"`
	Status    TripStatus `gorm:"null" json:"status"`

	Courier Courier `gorm:"foreignKey:CourierID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Order   Order   `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
