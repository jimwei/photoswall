package common

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	//database server
	DBSERVER string = "mysql"
	dbUser   string = "sa"
)

var (
	// db object
	DB gorm.DB
)

func init() {
	log.Println("init gorm.")
	var err error = nil
	dbPass := os.Getenv("DBA_PASS")
	DB, err = gorm.Open(DBSERVER, fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/photoswall", dbUser, dbPass))
	if err != nil {
		log.Fatal("initialize the dbtabase server fail.", err.Error())
	}
}
