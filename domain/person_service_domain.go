package domain

import "github.com/CharanDetDev/go-port-adapter/model"

type PersonService interface {
	GetPersonWithPersonID(personId int) (*model.Person, error)
	InsertPerson(newPerson *model.PersonRequest) error
	UpdatePerson(newPerson *model.PersonRequest) error
	DeletePerson(personID int) error
}
