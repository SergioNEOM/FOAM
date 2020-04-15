package database

import (
	"github.com/SergioNEOM/FOAM/models"
)

//NewUser - create new record in Users
func NewUser(udata models.User) (userid uint) { // return 0 == error!
	//	if err := g.DB.NewRecord(&udata).Error; err == nil {
	if err := Dbase.Create(&udata).Error; err != nil {
		return 0
	}
	//}
	return udata.ID
}

// GetUserByID get user info from DB
func GetUserByID(id uint) (*models.User, error) {
	u := &models.User{}
	Dbase.First(u, id)
	return u, nil // !!!
}

//UpdUser - update ???
func UpdUser(u *models.User) {
	Dbase.Save(u)
}

//DelUser delete by Id
func DelUser(userid uint) bool {
	u := &models.User{ID: userid}
	Dbase.First(&u, userid) //todo: а надо ли эту функцию вызывать?
	err := Dbase.Delete(&u).Error
	return (err != nil)
}
