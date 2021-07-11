package usecases

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sr-2020/gateway/app/adapters/storage"
	"github.com/sr-2020/gateway/app/domain"
	"strings"
)

type GenerateJwtInterface interface {
	Execute(GenerateJwtRequest) (GenerateJwtResponse, error)
}

type GenerateJwtRequest struct {
	Host   string
	Method string
	Scopes []string
	Tokens []domain.Payload
}

type GenerateJwtResponse struct {
	Result string
}

type GenerateJwt struct {
	Secret  string
	Storage storage.Storage
}

func (j *GenerateJwt) templGet(request GenerateJwtRequest, payload domain.Payload, token string) string {
	headers := fmt.Sprintf(`[Host: %s]
[Authorization: Bearer %s]`, request.Host, token)
	endpoints := ""
	for _, v := range request.Scopes {
		switch v {
		case "auth": endpoints = endpoints + "\n" + fmt.Sprintf(`/__auth__?modelId=%d auth`, payload.ModelId)
		case "model": endpoints = endpoints + "\n" + fmt.Sprintf(`/api/v1/models-manager/character/model?modelId=%d model`, payload.ModelId)
		case "billing": endpoints = endpoints + "\n" + fmt.Sprintf(`/api/v1/billing/sin?modelId=%d balance
/api/v1/billing/transfers?modelId=%d transfers
/api/v1/billing/rentas?modelId=%d rentas
/api/v1/billing/api/Scoring/info/getmyscoring?modelId=%d myscoring`, payload.ModelId, payload.ModelId, payload.ModelId, payload.ModelId)
		}
	}

	return headers + endpoints
}

func (j *GenerateJwt) templPost(request GenerateJwtRequest, payload domain.Payload, token string) string {
	headers := fmt.Sprintf(`[Host: %s]
[Content-Type: application/json]
[Authorization: Bearer %s]`, request.Host, token)
	endpoints := ""
	endpointTempl := `152 /api/v1/position/positions?modelId=%d&locationId=%d position`
	dataTempl := `{"beacons":[{"ssid":"F5:FD:E6:A1:69:CB","bssid":"F5:FD:E6:A1:69:CB","level":%d},{"ssid":"D8:28:B9:6B:CF:79","bssid":"D8:28:B9:6B:CF:79","level":-40}]}`
	for _, v := range request.Scopes {
		switch v {
		case "position": {
			endpoints = endpoints + "\n" + fmt.Sprintf(endpointTempl, payload.ModelId, 52)
			endpoints = endpoints + "\n" + fmt.Sprintf(dataTempl, -30)
		}
		case "position-change": {
			endpoints = endpoints + "\n" + fmt.Sprintf(endpointTempl, payload.ModelId, 52)
			endpoints = endpoints + "\n" + fmt.Sprintf(dataTempl, -30)
			endpoints = endpoints + "\n" + fmt.Sprintf(endpointTempl, payload.ModelId, 54)
			endpoints = endpoints + "\n" + fmt.Sprintf(dataTempl, -50)
		}
		}
	}

	return headers + endpoints
}

func (j *GenerateJwt) Execute(request GenerateJwtRequest) (GenerateJwtResponse, error) {
	var response GenerateJwtResponse

	tokens := make([]string, 0)
	for _, v := range request.Tokens {
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
			"auth": v.Auth,
			"modelId": v.ModelId,
			"exp": v.Exp,
		})

		tokenString, err := token.SignedString([]byte(j.Secret))
		if err != nil {
			return response, err
		}

		if request.Method == "post" {
			tokens = append(tokens, j.templPost(request, v, tokenString))
		} else {
			tokens = append(tokens, j.templGet(request, v, tokenString))
		}
	}

	response.Result = strings.Join(tokens, "\n\n")

	return response, nil
}
