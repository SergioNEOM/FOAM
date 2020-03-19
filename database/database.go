package database

import (
	"github.com/jinzhu/gorm"

	"github.com/SergioNEOM/FOAM/models"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// GDB wrap  gorm.DB - main database of app
type GDB struct {
	DB *gorm.DB
}

var Dbase *GDB

func init() {
	// init DB
	// dev mode only:
	Dbase, err := New("sqlite3", "./test.db")

	// release mode: 1) get conn params 2) Dbase, err := New(par1,par2)

	if err != nil {
		panic(err)
	}

	defer Dbase.Close()
}

// New create a new instance of gorm.DB, open it and return *GDB
func New(dialect, connstr string) (*GDB, error) {
	// в зависимости от диалекта проинициализировать параметры и открыть БД
	db, err := gorm.Open(dialect, connstr)
	if err != nil {
		return nil, err
	}
	// if dialect == ...
	//
	if dialect == "sqlite3" {
		// no concurrent connections
		//( https://github.com/mattn/go-sqlite3/issues/274 )
		db.DB().SetMaxOpenConns(1)
	}
	//
	//------------------------------------
	db.AutoMigrate(&models.User{}, &models.Stuff{})
	//------------------------------------

	return &GDB{DB: db}, nil
}

// Close closes DB
func (g *GDB) Close() {
	g.DB.Close()
}

func (g *GDB) AddStuff(s *models.Stuff) error {
	return g.DB.Create(s).Error
}
