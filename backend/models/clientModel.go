package models

import (
	"gorm.io/gorm"
	"time"
)

type Client struct {
	ID             uint           `json:"id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	ClientName     string         `json:"client_name" gorm:"unique"`
	ClientInitials string         `json:"client_initials" gorm:"unique"`
	WorkOrders     []WorkOrder    `json:"work_orders" gorm:"constraint:OnUpdate:CASCADE"`
}
