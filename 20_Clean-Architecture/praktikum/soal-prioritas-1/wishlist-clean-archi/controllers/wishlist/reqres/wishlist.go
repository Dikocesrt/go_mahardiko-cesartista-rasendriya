package reqres

import (
	"time"

	"gorm.io/gorm"
)

type Wishlist struct {
	ID         uint `json:"id" form:"id"`
	CreatedAt  time.Time `json:"created_at" form:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" form:"deleted_at"`
	Title      string `json:"title" form:"title"`
	IsAchieved bool `json:"is_achieved" form:"is_achieved"`
}