package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Debug(c echo.Context) error {
	body := echo.Map{}
	if err := c.Bind(&body); err != nil {
		return err
	}

	data := echo.Map{}
	data["headers"] = c.Request().Header
	data["body"] = body

	return c.JSON(http.StatusOK, data)
}
