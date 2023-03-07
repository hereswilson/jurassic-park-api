package models

import (
	"gorm.io/gorm"
)

type Species struct {
	gorm.Model
	Species string `json:"species" gorm:"type:text; not null"`
	Diet    string `json:"diet" gorm:"type:text; not null"`
}
