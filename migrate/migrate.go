package main

import (
	"fmt"
	"go-clean-architecture/db"
	"go-clean-architecture/entity"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated!")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&entity.User{}, &entity.Task{})
}
