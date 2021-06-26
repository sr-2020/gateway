package usecases

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sr-2020/gateway/app/adapters/storage"
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
	Secret  string
	Storage storage.Storage
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
		if exp, ok := claims["exp"].(float64); ok {
			payload.Exp = int64(exp)
		}
	}

	//if payload.Auth == domain.RolePlayer {
	//	key := strconv.Itoa(payload.ModelId)
	//	if !j.Storage.Check(key, payload.Exp) {
	//		return response, domain.ErrMultiLogin
	//	}
	//}

	response.Payload = payload

	return response, nil
}
