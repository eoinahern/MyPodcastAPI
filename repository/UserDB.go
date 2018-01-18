package repository

import (
	"fmt"
	"my_podcast_api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserDB struct {
	*gorm.DB
}

func (DB *UserDB) CheckExist(email string) bool {

	var count int = 0
	DB.Model(&models.User{}).Where("user_name = ?", email).Count(&count)

	if count >= 1 {
		return true
	}

	return false
}

func (DB *UserDB) ValidatePasswordAndUser(email string, password string) bool {

	str := fmt.Sprintf("passed in : ? , ? ", email, password)
	fmt.Println(str)

	var user models.User
	DB.Where("user_name = ? AND password = ?", email, password).First(&user)

	fmt.Println(user.UserName)

	if user.UserName == email {
		return true
	} else {
		return false
	}
}

func (DB *UserDB) Insert(user *models.User) {
	DB.Save(user)
}

func (DB *UserDB) GetUser(email string) models.User {

	var user models.User
	DB.Where("user_name = ?", email).First(&user)
	return user
}

func (DB *UserDB) GetAll() {

}

func (DB *UserDB) update(email string) {

}

func (DB *UserDB) delete(email string) bool {

	//delete user if worked return true
	return true

}
