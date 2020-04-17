package database

import (
	"github.com/SergioNEOM/FOAM/models"
)

// GetStuffList get all stuff from DB
//
func GetStuffList() *[]models.Stuff {

	//s := &[]models.Stuff{}

	// !!! test & dev mode only:
	// init test stuff
	s := &[]models.Stuff{
		{ShortName: "testrec1", Description: "12345678901234567890"},
		{ShortName: "testrec2", Description: "0-1-2-3-4-5-6-7-8-9-0-1-2-3-4-5-6-7-8-9-"},
	}
	//
	// release mode - get stuff list from DB
	// g.
	return s // !!!
}

// AddStuff - добавить материал в БД
func AddStuff(s *models.Stuff) error {
	return DB.Create(s).Error
}
