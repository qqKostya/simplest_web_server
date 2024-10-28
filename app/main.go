package main

import (
	"github.com/labstack/echo"
	"simplest_web_server/app/handlers"
)

func main() {
	e := echo.New()
	e.GET("/stat.js", handlers.Stat)
	e.GET("/counter.js", handlers.Counter)
	e.Start(":8080")
}
