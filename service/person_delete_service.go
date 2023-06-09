package service

import (
	"fmt"

	"github.com/CharanDetDev/go-port-adapter/util/logg"
)

func (service *personService) DeletePerson(personID int) error {

	_, err := service.PersonRepo.GetPersonWithPersonID(personID)
	if err != nil {
		logg.Printlogger("DELETE failed", "not found person id", fmt.Sprintf("%v | %v", err.Error(), logg.GetCallerPathNameFileNameLineNumber()))
		return fmt.Errorf("gorm delete failed, not found person id")
	}

	err = service.PersonRepo.DeletePerson(personID)
	if err != nil {
		logg.Printlogger("DELETE failed", "", fmt.Sprintf("%v | %v", err.Error(), logg.GetCallerPathNameFileNameLineNumber()))
		return fmt.Errorf("gorm delete failed")
	}

	return nil

}
