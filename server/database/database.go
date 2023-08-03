package database

import (
	"fmt"
	"os"
	"typing/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Open() {

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", host, user, pass, name)

	db, errDb := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if errDb != nil {
		panic(errDb.Error())
	}

	if errMigrate := db.AutoMigrate(&entities.User{}, &entities.Score{}); errMigrate != nil {
		panic(errMigrate)
	}
	Db = db
}

func Close() {
	dbSql, errDb := Db.DB()

	if errDb != nil {
		panic(errDb.Error())
	}
	if errClose := dbSql.Close(); errClose != nil {
		panic(errClose.Error())
	}
}
