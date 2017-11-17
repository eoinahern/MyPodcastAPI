package repository

import "my_podcast_api/models"

type UserDB struct {
}

func (DB *UserDB) CheckExist(email string) bool {
	return false
}

func (DB *UserDB) Insert(user *models.User) {

	//inset a row into the database table for users

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
