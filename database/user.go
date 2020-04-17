package database

import (
	"errors"
	"fmt"

	"github.com/SergioNEOM/FOAM/models"
)

// Интерфейс для работы с таблицей пользователей в БД
type UserMaker interface {
	AddUser() uint
	GetUser()
	GetUserByID(uint) (*models.User, error)
	GetAllUser() ([]models.User, error)
	UpdUser()
	DelUser(uint)
}

//NewUser - create new record in Users
func NewUser(to, lo, pa, ro, na string) (userid uint) { // return 0 == error!
	u := &models.User{
		Token: to,
		Login: lo,
		Pass:  pa,
		Role:  ro,
		Name:  na,
	}
	return NewUserRec(u)
}

//NewUserRec - create new record in Users
func NewUserRec(udata *models.User) (userid uint) { // return 0 == error!
	fmt.Println("[NewUserRec] -- begin")
	//	if err := g.DB.NewRecord(&udata).Error; err == nil {
	db := GetDB()
	if db == nil {
		fmt.Println("[NewUserRec] -- Dbase is nil")
		return 0
	}
	if err := db.Create(udata).Error; err != nil {
		fmt.Println("[NewUserRec] -- error")
		return 0
	}
	//}
	fmt.Println("[NewUserRec] -- end: ", udata.ID)
	return udata.ID
}

// GetUserByID get user info from DB
func GetUserByID(id uint) (*models.User, error) {
	u := &models.User{}
	if db := GetDB(); db != nil {
		db.First(u, id)
		return u, nil // !!!
	}
	return nil, errors.New("[GetUserByID] error")
}

//UpdUser - update ???
func UpdUser(u *models.User) {
	//todo: if u == nil {}
	if db := GetDB(); db != nil {
		db.Save(u)
	}
}

//DelUser delete by Id
func DelUser(userid uint) bool {
	u := &models.User{ID: userid}
	if db := GetDB(); db != nil {
		db.First(&u, userid) //todo: а надо ли эту функцию вызывать?
		err := db.Delete(&u).Error
		return (err != nil)
	}
	return false
}

func GetAllUsers() ([]models.User, error) {
	ua := []models.User{}
	var err error
	if db := GetDB(); db != nil {
		err = db.Find(&ua).Error
	}
	return ua, err
}
