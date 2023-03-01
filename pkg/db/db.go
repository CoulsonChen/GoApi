package db

import (
	"fmt"

	"github.com/CoulsonChen/GoApi/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBProvider(dbconfig config.DbConfig) (db *gorm.DB) {

	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbconfig.Host,
		dbconfig.User,
		dbconfig.Password,
		dbconfig.Dbname,
		dbconfig.Port,
		dbconfig.Sslmode)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		// log.Fatal(err)
		panic(err)
	}
	// sqlConn, err := db.DB()
	// if err != nil {
	// 	panic(err)
	// }
	// defer sqlConn.Close()

	return
}
