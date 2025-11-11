package entity

import (
	"time"
)

type Article struct {
	ID          int64     `json:"id" xorm:"pk autoincr BIGINT(20) id"`
	Title       string    `json:"title" xorm:"not null default '' VARCHAR(255) title"`
	Slug        string    `json:"slug" xorm:"not null default '' VARCHAR(255) unique slug"`
	Content     string    `json:"content" xorm:"LONGTEXT content"`
	Status      string    `json:"status" xorm:"not null default 'draft' VARCHAR(20) status"` // draft, published, archived
	UserID      int64     `json:"user_id" xorm:"not null default 0 BIGINT(20) user_id"`
	MenuID      int64     `json:"menu_id" xorm:"not null default 0 BIGINT(20) menu_id"`
	FeaturedImg string    `json:"featured_img" xorm:"VARCHAR(500) featured_image"`
	CreatedAt   time.Time `json:"created_at" xorm:"created TIMESTAMP created_at"`
	UpdatedAt   time.Time `json:"updated_at" xorm:"updated TIMESTAMP updated_at"`
	PublishedAt time.Time `json:"published_at" xorm:"TIMESTAMP published_at"`
}

// TableName
func (Article) TableName() string {
	return "articles"
}
