package entities

import "gorm.io/gorm"

type Transaksi struct {
	gorm.Model
	Nota   string
	Alamat string
}
