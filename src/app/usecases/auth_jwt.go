package usecases

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sr-2020/gateway/app/domain"
)

type JwtInterface interface {
	Execute(JwtRequest) (JwtResponse, error)
}

type JwtRequest struct {
	Token  string
}

type JwtResponse struct {
	Payload domain.Payload
}

type Jwt struct {
	Secret string
}

func (j *Jwt) Execute(request JwtRequest) (JwtResponse, error) {
	var response JwtResponse
	var payload domain.Payload

	token, err := jwt.Parse(request.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return response, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.Secret), nil
	})

	if err != nil {
		return response, fmt.Errorf("JWT Token is invalid: %s", err.Error())
	}

	if token == nil {
		return response, fmt.Errorf("JWT Token is invalid: nil")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if modelId, ok := claims["modelId"].(float64); ok {
			payload.ModelId = int(modelId)
		}
		if auth, ok := claims["auth"].(string); ok {
			payload.Auth = auth
		}
	}

	response.Payload = payload

	return response, nil
}
