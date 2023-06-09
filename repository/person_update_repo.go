package repository

import (
	"github.com/CharanDetDev/go-port-adapter/util/database"

	"github.com/CharanDetDev/go-port-adapter/model"
)

func (repo *personRepo) UpdatePerson(newPerson *model.Person) error {

	result := database.Conn.Model(&newPerson).Where("PersonID = ?", newPerson.PersonID).Updates(newPerson)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
