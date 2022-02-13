package database

import (
	"planwey/model"
)

func Migrate() error {
	err := Db.AutoMigrate(&model.User{}, &model.Timeslot{}, &model.Task{})
	return err
}
