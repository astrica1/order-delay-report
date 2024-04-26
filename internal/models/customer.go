package models

import (
	"errors"
	"strings"

	"github.com/astrica1/order-delay-report/pkg/validator"
	"gorm.io/gorm"
)

type Customer struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Username    string `gorm:"unique;not null" json:"username"`
	Password    string `gorm:"type:varchar(32)" json:"password"` //store passwords as MD5 hash
	FullName    string `gorm:"null" json:"full_name"`
	PhoneNumber string `gorm:"type:varchar(10);null" json:"phone_number"`
}

func (c *Customer) isValidUsername(db *gorm.DB, username string) error {
	if err := validator.ValidateUsername(username); err != nil {
		err := errors.New("username can only contain lowercase letters and underscores: " + username)
		db.AddError(err)
		return err
	}
	return nil
}

func (c *Customer) BeforeCreate(db *gorm.DB) (err error) {
	c.Username = strings.ToLower(c.Username)
	err = c.isValidUsername(db, c.Username)
	return err
}
