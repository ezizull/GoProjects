package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite
)

// Lead exported
type Lead struct {
	gorm.Model
	Name 	string	`json:"name"`
	Company string	`json:"company"`
	Email 	string	`json:"email"`
	Phone 	int 	`json:"phone"`
}