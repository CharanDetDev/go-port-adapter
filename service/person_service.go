package service

import "github.com/CharanDetDev/go-port-adapter/domain"

type personService struct {
	PersonRepo domain.PersonRepo
}

func NewPersonService(personRepo domain.PersonRepo) domain.PersonService {
	return &personService{
		PersonRepo: personRepo,
	}
}
