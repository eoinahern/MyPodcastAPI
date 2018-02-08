package repository

import (
	"fmt"
	"log"
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

func (DB *UserDB) ValidateUserPlusRegToken(email string, regToken string) bool {

	var count int = 0
	DB.Model(&models.User{}).Where("user_name = ? AND reg_token = ?", email, regToken).Count(&count)

	if count == 1 {
		return true
	}

	return false
}

func (DB *UserDB) SetVerified(username string, token string) {

	var user models.User
	DB.Where("user_name = ? AND reg_token = ?", username, token).First(&user)
	user.Verified = true
	db := DB.Save(&user)

	if db.Error != nil {
		log.Println(db.Error)
	}

}

func (DB *UserDB) ValidatePasswordAndUser(email string, password string) bool {

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
