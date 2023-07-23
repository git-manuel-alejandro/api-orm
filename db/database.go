package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

var Database = func() (db *gorm.DB) {
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
		return
	}
	// dns := "root:'':9K3@tcp(localhost:3306)/golang_orm?charset=utf8mb4&parseTime=True&loc=Local"

	dns := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_SERVER") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME")

	if db, err := gorm.Open(mysql.Open(dns), &gorm.Config{}); err != nil {
		fmt.Println("error de conexión")
		panic(err)
	} else {
		fmt.Println("conexión exitosa")
		return db
	}

}()
