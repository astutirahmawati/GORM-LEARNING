package datastore

import (
	"gormm/entities"

	"fmt"

	"gorm.io/gorm"
)

type UserDB struct {
	Db *gorm.DB
}

func (u *UserDB) GetAllDataUser() ([]entities.User, error) {
	res := []entities.User{}
	if err := u.Db.Table("User").Where("Name LIKE ?", "%J%").Find(&res).Error; err != nil {
		fmt.Println("Terjadi kesalahan", err)
		return []entities.User{}, err
	}
	return res, nil
}
