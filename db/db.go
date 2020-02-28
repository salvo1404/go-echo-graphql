package db

import (
	"log"

	"github.com/go-sql-driver/mysql" // Mysql driver
	"github.com/jinzhu/gorm"
)

func Connect() *gorm.DB {
	DBMS := "mysql"
	mySQLConfig := &mysql.Config{
		User:                 "go-echo-grapql",
		Passwd:               "go-echo-grapql",
		Net:                  "tcp",
		Addr:                 "db:3308",
		DBName:               "go-echo-grapql",
		AllowNativePasswords: true,
		Params: map[string]string{
			"parseTime": "true",
		},
	}

	db, err := gorm.Open(DBMS, mySQLConfig.FormatDSN())
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
