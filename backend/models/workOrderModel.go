package models

import (
	"gorm.io/gorm"
	"time"
)

type WorkOrder struct {
	ID            uint           `json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
	WorkOrderName string         `json:"work_order_name"`
	ClientPO      string         `json:"client_po"`
	Requestor     string         `json:"requestor"`
	ProjectName   string         `json:"project_name"`
	WorkOrderDate time.Time      `json:"work_order_date"`
}
