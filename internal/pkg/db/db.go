package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBProvider() (db *gorm.DB) {
	dbURL := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlConn, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlConn.Close()

	return
}
