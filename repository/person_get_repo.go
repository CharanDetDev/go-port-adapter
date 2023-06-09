package repository

import (
	"github.com/CharanDetDev/go-port-adapter/util/database"

	"github.com/CharanDetDev/go-port-adapter/model"
)

func (repo *personRepo) GetPersonWithPersonID(personId int) (*model.Person, error) {

	var person model.Person
	err := database.Conn.First(&person, personId).Error
	if err != nil {
		return &person, err
	}

	return &person, nil
}
