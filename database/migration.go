package database

import (
	"fmt"
	"test/models"
	"test/pkg/mysql"
)

func RunMigrate() {
	if err := mysql.DB.AutoMigrate(
		&models.Todo{},
		&models.Activity{},
	); err != nil {
		fmt.Println("err")
		panic("Migration Failed")
	}
	fmt.Println("Migration Success")
}
