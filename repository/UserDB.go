package repository

import (
	"fmt"
	"my_podcast_api/models"

	"github.com/jinzhu/gorm"
)

type UserDB struct {
	*gorm.DB
}

func (DB *UserDB) CheckExist(email string) bool {

	var count int = 0
	DB.Model(&models.User{}).Where("user_name = ?", email).Count(&count)

	str := fmt.Sprintf("count : %d", count)
	fmt.Println(str)

	if count >= 1 {
		return true
	}

	return false
}

func (DB *UserDB) Insert(user *models.User) {

	//inset a row into the database table for users
	DB.Save(user)

}

func (DB *UserDB) GetItem(email string) *models.User {

	return &models.User{}
}

func (DB *UserDB) GetAll() {

}

func (DB *UserDB) update(email string) {

}

func (DB *UserDB) delete(email string) bool {

	//delete user if worked return true
	return true

}
