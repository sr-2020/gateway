package usecases

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/sr-2020/gateway/app/domain"
	"reflect"
	"testing"
	"time"
)

const (
	jwtSecret = "test"
)

func jwtToken(id int, secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub":         "37445",
		"auth":        "ROLE_PLAYER",
		"modelId":     id,
		"characterId": 64,
		"exp":         time.Now().Add(1 * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}

	return tokenString
}

func TestJwt_Execute(t *testing.T) {
	type args struct {
		request JwtRequest
	}
	tests := []struct {
		name    string
		jwt     *Jwt
		args    args
		want    JwtResponse
		wantErr bool
	}{
		{
			name: "Success for 1",
			jwt: &Jwt{
				Secret: jwtSecret,
			},
			args: args{
				request: JwtRequest{
					Token: jwtToken(1, jwtSecret),
				},
			},
			want: JwtResponse{
				Payload: domain.Payload{
					Sub:     "",
					Auth:    "",
					ModelId: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "Error secret",
			jwt: &Jwt{
				Secret: "test2",
			},
			args: args{
				request: JwtRequest{
					Token: jwtToken(1, jwtSecret),
				},
			},
			want: JwtResponse{
				Payload: domain.Payload{
					Sub:     "",
					Auth:    "",
					ModelId: 0,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.jwt.Execute(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Jwt.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Jwt.Handle() = %v, want %v", got, tt.want)
			}
		})
	}
}
