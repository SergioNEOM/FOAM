package config

import (
	"fmt"
	"os"

	"github.com/sasbury/mini"
)

type foamConf struct {
	LogFile   string
	DBDialect string
	DBConnStr string
}

//Conf - global config values
var Conf *foamConf

func init() {
	cfile := "-"
	if len(os.Args) > 1 {
		cfile = os.Args[1]
	}
	Conf = loadConfig(cfile)
}

//loadConfig - get config values from *.ini file
func loadConfig(filename string) *foamConf {
	// default values
	c := &foamConf{LogFile: "", DBDialect: "sqlite3", DBConnStr: "./foam_default.db"}
	//todo: проверить работу с ini-файлом
	cf, err := mini.LoadConfiguration(filename)
	if err != nil {
		fmt.Println("--No config file --  default values will be set")
	} else {
		c.LogFile = cf.StringFromSection("common", "log_file", "")
		c.DBDialect = cf.StringFromSection("database", "dialect", "sqlite3")
		c.DBConnStr = cf.StringFromSection("database", "connection_string", "./foam_default.db")
	}
	return c
}
