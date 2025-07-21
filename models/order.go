package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID     string `gorm:"primaryKey;column:order_id" json:"order_id"`
	UserID uint   `json:"user_id"`
	Amount int    `json:"amount"`
	Status string `json:"status"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID == "" {
		o.ID = uuid.NewString()
	}
	return
}
