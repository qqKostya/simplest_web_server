package main

import (
	"fmt"
	ds "simplest_web_server/app/datastore"
	"simplest_web_server/app/handlers"
	"time"

	"github.com/labstack/echo"
)

func main() {
	ds := ds.Datastore{}

	go printHits(&ds)

	e := echo.New()
	e.GET("/stat.js", handlers.Stat(&ds))
	e.GET("/counter.js", handlers.Counter(&ds))
	e.Start(":8080")
}

func printHits(ds *ds.Datastore) {
	for {
		time.Sleep(time.Second)
		fmt.Printf("Total visits %v\n", ds.TotalHits())
	}
}
