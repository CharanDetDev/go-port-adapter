package main

import (
	"fmt"

	"github.com/CharanDetDev/go-port-adapter/api/route"
	"github.com/CharanDetDev/go-port-adapter/util/cache"
	"github.com/CharanDetDev/go-port-adapter/util/config"
	"github.com/CharanDetDev/go-port-adapter/util/database"
	"github.com/CharanDetDev/go-port-adapter/util/logg"
	"github.com/gofiber/fiber/v2"
)

func init() {

	isConfig := config.ConfigInit()
	isDatabase := database.InitDatabase()
	isCache := cache.InitCache()
	if isConfig && isDatabase && isCache {
		logg.Printlogger_Variadic("\t ***** Initail :: Configuration & Database & Redis :: SUCCESS **** ", "Results", *database.Conn, cache.RedisCaching.RedisClient)
	} else {
		logg.Printlogger_Variadic("\t ***** Initail :: Configuration & Database & Redis :: ERROR **** ", "Results", *database.Conn, cache.RedisCaching.RedisClient, logg.GetCallerPathNameFileNameLineNumber())
		panic(fmt.Errorf("initail configuration error"))
	}

}

func main() {
	defer database.ConnectionClose()

	app := fiber.New()
	router := route.NewRoute()
	router.InitRoute(app)

	app.Listen(config.Env.API_PORT)
}
