package database

import (
	"log"

	"github.com/jinzhu/gorm"

	"github.com/SergioNEOM/FOAM/config"
	"github.com/SergioNEOM/FOAM/models"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// GDB wrap  gorm.DB - main database of app
/*type GDB struct {
	DB *gorm.DB
}
*/

//
var Dbase *gorm.DB

func init() {
	// init DB
	// dev mode only:
	log.Printf("[database init] %v, %v\n", config.Conf.DBDialect, config.Conf.DBConnStr)
	// в зависимости от диалекта проинициализировать параметры и открыть БД
	Dbase, err := gorm.Open(config.Conf.DBDialect, config.Conf.DBConnStr)

	// release mode: 1) get conn params 2) Dbase, err := New(par1,par2)

	if err != nil || Dbase == nil {
		log.Fatal("Dbase not opened")
		panic(err)
	}
	defer Dbase.Close()
	//
	if config.Conf.DBDialect == "sqlite3" {
		// no concurrent connections
		//( https://github.com/mattn/go-sqlite3/issues/274 )
		Dbase.DB().SetMaxOpenConns(1)
	}
	//
	//------------------------------------
	Dbase.AutoMigrate(&models.User{}, &models.Stuff{})
}
