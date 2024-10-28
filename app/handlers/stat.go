package handlers

import (
    "github.com/labstack/echo"
    "net/http"
)

// Stat handler
func Stat(c echo.Context) error {
    return c.String(http.StatusOK, `{ "total": 0 }`)
}