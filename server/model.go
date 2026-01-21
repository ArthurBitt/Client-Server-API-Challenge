package server

import "gorm.io/gorm"

type Cotacao struct {
	gorm.Model
	Bid string `gorm:"type:varchar(20)"`
}
