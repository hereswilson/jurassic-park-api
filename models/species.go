package models

import "gorm.io/gorm"

type Species struct {
	gorm.Model
	Species string `json:"species"`
	Diet    string `json:"diet"`
}
