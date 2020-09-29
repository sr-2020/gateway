package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/sr-2020/gateway/app/usecases"
	"net/http"
	"strconv"
	"strings"
)

type Auth struct {
	UseCase     usecases.JwtInterface
	UseCaseData usecases.DataInterface
}

func (a Auth) Handler(c echo.Context) error {
	token := ""
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader != "" {
		auth := strings.Split(c.Request().Header.Get("Authorization"), " ")
		if len(auth) < 2 {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Invalid token header",
			})
		}
		token = auth[1]
	} else {
		tokenCookie, err := c.Request().Cookie("Authorization")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Invalid token cookie:" + err.Error(),
			})
		}

		token = tokenCookie.Value
	}

	request := usecases.JwtRequest{Token: token}
	response, err := a.UseCase.Execute(request)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
	}

	data := c.QueryParam("data")
	if data != "" {
		requestData := usecases.DataRequest{
			Id: response.Payload.ModelId,
			Scopes: []string{data},
		}
		responseData, err := a.UseCaseData.Execute(requestData)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": err.Error(),
			})
		}

		userData, err := json.Marshal(responseData.Data)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		c.Response().Header().Set("X-User-Data", string(userData))
	}

	c.Response().Header().Set("X-User-Id", strconv.Itoa(response.Payload.ModelId))

	return c.JSON(http.StatusOK, struct {}{})
}
