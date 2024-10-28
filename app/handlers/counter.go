package handlers

import (
	"github.com/labstack/echo"
	"net/http"
	ds "simplest_web_server/app/datastore"
)

// Counter handler
func Counter(ds *ds.Datastore) echo.HandlerFunc {
	return func(c echo.Context) error {
		ds.RegisterHit()
		return c.String(http.StatusOK, "ok")
	}
}
