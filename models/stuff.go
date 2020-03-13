package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

/*
  Stuff  - обобщенная единица хранения учебных материалов в БД
*/
type Stuff struct {
	gorm.Model
	// краткое наименование (заголовок)
	ShortName string `gorm:"type:varchar(100)"`
	// Описание
	Description string `gorm:"type:varchar(255)"`
	// тип материалов
	//
	// ссылка на файл
	//
}
