package models

import (
	"errors"
	"strings"

	"github.com/astrica1/order-delay-report/pkg/validator"
	"gorm.io/gorm"
)

type Agent struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"type:varchar(32)" json:"password"` //store passwords as MD5 hash
	IsActive bool   `gorm:"default:true" json:"is_active"`
}

func (a *Agent) isValidUsername(db *gorm.DB, username string) error {
	if err := validator.ValidateUsername(username); err != nil {
		err := errors.New("username can only contain lowercase letters and underscores: " + username)
		db.AddError(err)
		return err
	}
	return nil
}

func (a *Agent) BeforeCreate(db *gorm.DB) (err error) {
	a.Username = strings.ToLower(a.Username)
	err = a.isValidUsername(db, a.Username)
	return err
}
