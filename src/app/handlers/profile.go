package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sr-2020/gateway/app/usecases"
	"net/http"
	"strconv"
)

type Profile struct {
	UseCase     usecases.JwtInterface
	UseCaseData usecases.DataInterface
}

type PositionLocation struct {
	Id      int                    `json:"id"`
	Label   string                 `json:"label"`
}

type PositionUser struct {
	Id         int              `json:"id"`
	Location   PositionLocation `json:"location"`
}

func (a Profile) Handler(c echo.Context) error {
	userIdHeader := c.Request().Header.Get("X-User-Id")
	if userIdHeader == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Empty auth header",
		})
	}

	userId, err := strconv.Atoi(userIdHeader)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	res := make([]PositionUser, 0)

	data := c.QueryParam("data")
	if data != "" {
		requestData := usecases.DataRequest{
			Id: userId,
			Scopes: []string{data},
		}
		responseData, err := a.UseCaseData.Execute(requestData)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": err.Error(),
			})
		}

		res = append(res, PositionUser{
			Id: userId,
			Location: PositionLocation{
				Id: responseData.Location.Id,
				Label: responseData.Location.Label,
			},
		})

		return c.JSON(http.StatusOK, res)
	}

	return c.JSON(http.StatusOK, res)
}
