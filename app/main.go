package main

import (
	"github.com/labstack/echo"
	"simplest_web_server/app/datastore"
	"simplest_web_server/app/handlers"
)

func main() {
	ds := datastore.Datastore{}

	e := echo.New()
	e.GET("/stat.js", handlers.Stat(&ds))
	e.GET("/counter.js", handlers.Counter(&ds))
	e.Start(":8080")
}
