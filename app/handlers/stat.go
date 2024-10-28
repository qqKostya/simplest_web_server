package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	ds "simplest_web_server/app/datastore"
)

func Stat(ds *ds.Datastore) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf(`{ "total": %d }`, ds.TotalHits()))
	}
}
