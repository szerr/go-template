package main

import (
	"go-template/internal/model"
	"gorm.io/gorm"
)

func main() {
	run, cleanup, err := wireApp()
	if err != nil {
		panic(err)
	}
	defer cleanup()
	err = run()
	if err != nil {
		panic(err)
	}
}

func newApp(modelLi model.AllModel, db *gorm.DB) func() error {
	return func() error { return db.AutoMigrate(modelLi...) }
}
