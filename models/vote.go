package models

import "github.com/jinzhu/gorm"

// Vote estructura que define los datos de nuestro modelo Vote
type Vote struct {
	gorm.Model
	CommentID uint `json:"commentId" gorm:"not null"`
	UserID    uint `json:"userId" gorm:"not null"`
	Value     bool `json:"value" gorm:"not null"`
}
