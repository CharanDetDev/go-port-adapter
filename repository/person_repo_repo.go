package repository

import "github.com/CharanDetDev/go-port-adapter/domain"

type personRepo struct{}

func NewPersonRepo() domain.PersonRepo {
	return &personRepo{}
}
