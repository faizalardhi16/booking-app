package models

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Model struct {
	Model interface{}
}

func listRegistry() []Model {
	return []Model{
		{Model: Todo{}},
		{Model: ImageDestination{}},
		{Model: Rating{}},
		{Model: Destination{}},
		{Model: User{}},
		{Model: Avatar{}},
	}
}

func RegisterDB(db *gorm.DB) {

	for _, t := range listRegistry() {
		err := db.Debug().AutoMigrate(t.Model)

		if err != nil {
			log.Fatal(err.Error())
		}
	}

	fmt.Println("Success for migrating database")
}
