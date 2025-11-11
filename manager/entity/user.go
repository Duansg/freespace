package entity

import (
	"time"
)

type User struct {
	ID        int64     `json:"id" xorm:"pk autoincr BIGINT(20) id"`
	Name      string    `json:"name" xorm:"not null default '' VARCHAR(255) Name"`
	Email     string    `json:"email" xorm:"not null default '' VARCHAR(255) Email"`
	Password  string    `json:"password" xorm:"not null default '' VARCHAR(255) Password"`
	Status    string    `json:"status" xorm:"not null default 'draft' VARCHAR(20) Status"` // draft, published, archived
	Img       string    `json:"img" xorm:"VARCHAR(500) Img"`
	ExtJson   string    `json:"extJson" xorm:"VARCHAR(500) ExtJson"`
	CreatedAt time.Time `json:"created_at" xorm:"created TIMESTAMP CreatedAt"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated TIMESTAMP UpdatedAt"`
}

// TableName
func (User) TableName() string {
	return "Users"
}
