package entities

import "gorm.io/gorm"

type Barang struct {
	gorm.Model
	Nama string
}
