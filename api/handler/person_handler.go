package handler

import "github.com/CharanDetDev/go-port-adapter/domain"

type personHandler struct {
	PersonService domain.PersonService
}

func NewPersonHandler(PersonService domain.PersonService) domain.PersonHandler {
	return &personHandler{
		PersonService,
	}
}
