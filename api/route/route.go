package route

import (
	"github.com/CharanDetDev/go-port-adapter/api/handler"
	"github.com/CharanDetDev/go-port-adapter/domain"
	"github.com/CharanDetDev/go-port-adapter/repository"
	"github.com/CharanDetDev/go-port-adapter/service"
	"github.com/gofiber/fiber/v2"
)

type (
	route struct {
		PersonHandler domain.PersonHandler
	}

	Route interface {
		InitRoute(app *fiber.App)
		// InitRouteGroup(app fiber.Router)
	}
)

func NewRoute() Route {

	newPersonRepo := repository.NewPersonRepo()
	newPersonService := service.NewPersonService(newPersonRepo)
	newPersonHandle := handler.NewPersonHandler(newPersonService)

	return &route{
		newPersonHandle,
	}
}

func (route *route) InitRoute(app *fiber.App) {

	personGroup := app.Group("/person")
	route.personGroup(personGroup)

}

// func (route *route) InitRouteGroup(appGroup fiber.Router) {

// 	personGroup := appGroup.Group("/person")
// 	route.personGroup(personGroup)

// }
