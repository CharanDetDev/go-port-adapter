package handler

import (
	"github.com/CharanDetDev/go-port-adapter/model"
	"github.com/CharanDetDev/go-port-adapter/util/caller"
	"github.com/CharanDetDev/go-port-adapter/util/message"
	"github.com/CharanDetDev/go-port-adapter/util/validation"

	"github.com/gofiber/fiber/v2"
)

func (personHandler *personHandler) UpdatePerson(c *fiber.Ctx) error {

	var newPerson model.PersonRequest
	_, err := validation.Struct(c, &newPerson, message.InvalidParam)
	if err != nil {
		return caller.BadRequest(c, message.InvalidParam, err.Error())
	}

	if err := personHandler.PersonService.UpdatePerson(&newPerson); err != nil {
		return caller.InternalServerError(c, message.InvalidParam, err.Error())
	}

	return caller.Success(c, message.Success, message.Get(c, message.Success), nil)

}
