package repository

import (
	"github.com/CharanDetDev/go-port-adapter/util/database"

	"github.com/CharanDetDev/go-port-adapter/model"
)

func (repo *personRepo) InsertPerson(newPerson *model.Person) error {

	result := database.Conn.Create(newPerson)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
