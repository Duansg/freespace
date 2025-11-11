package entity

import (
	"time"
)

type Menu struct {
	ID        int64     `json:"id" xorm:"pk autoincr BIGINT(20) id"`
	Link      string    `json:"link" xorm:"not null default '' VARCHAR(255) Link"`
	Content   string    `json:"content" xorm:"not null default '' VARCHAR(50) Content"`
	Status    string    `json:"status" xorm:"not null default 'draft' VARCHAR(20) status"` // active
	Icon      string    `json:"icon" xorm:"not null default '' VARCHAR(50) Icon"`
	OrderNum  int64     `json:"order_num" xorm:"not null default 0 BIGINT(20) order_num"`
	CreatedAt time.Time `json:"created_at" xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated TIMESTAMP updated_at"`
}

// TableName
func (Menu) TableName() string {
	return "Menus"
}
