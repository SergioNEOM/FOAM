package database

import (
	"log"
	"sync"

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

// singletone !
var DB *gorm.DB

//  safety singletone maker
var once sync.Once

func init() {
	// init DB
	// dev mode only:
	log.Printf("[database init] %v, %v\n", config.Conf.DBDialect, config.Conf.DBConnStr)
	DB = GetDB()
	// в зависимости от диалекта проинициализировать параметры и открыть БД

	log.Printf("[database meta]\n %+v \n", DB)
	//defer Dbase.Close()
	//
	if config.Conf.DBDialect == "sqlite3" {
		// no concurrent connections
		//( https://github.com/mattn/go-sqlite3/issues/274 )
		DB.DB().SetMaxOpenConns(1)
	}
	//
	//------------------------------------
	DB.AutoMigrate(&models.User{}, &models.Stuff{})
	//
	//------------------------------------
	// default record
	//todo: исключить дубли!!!
	NewUser("", "admin", "admin", "admin", "учетная запись администратора по умолчанию")
}

//GetDB возвращает указатель на уникальное соединение с БД
func GetDB() *gorm.DB {
	once.Do(func() {
		// обеспечим уникальность (единственность) указателя на БД в пределах приложения - singletone
		db, err := gorm.Open(config.Conf.DBDialect, config.Conf.DBConnStr)
		if err != nil || db == nil {
			log.Fatal("Dbase not opened")
			return
		}
		DB = db
	})
	return DB
}
