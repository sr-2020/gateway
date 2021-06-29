package usecases

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/sr-2020/gateway/app/adapters/storage"
	"github.com/sr-2020/gateway/app/domain"
	"reflect"
	"testing"
	"time"
)

const (
	jwtSecret = "test"
)

func jwtToken(id int, auth string, exp int64, secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"auth":        auth,
		"modelId":     id,
		"characterId": 64,
		"exp":         exp,
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}

	return tokenString
}

func TestJwt_Execute(t *testing.T) {
	expOld := time.Now().Add(5 * time.Minute).Unix()
	expNew := time.Now().Add(10 * time.Minute).Unix()
	jwtPlayerOld := jwtToken(1, domain.RolePlayer, expOld, jwtSecret)
	jwtPlayerNew := jwtToken(1, domain.RolePlayer, expNew, jwtSecret)
	jwtMasterOld := jwtToken(2, domain.RoleMaster, expOld, jwtSecret)
	jwtMasterNew := jwtToken(2, domain.RoleMaster, expNew, jwtSecret)

	mockStorage := storage.NewMock(map[string][]string{
		"1": { jwtPlayerOld, jwtPlayerNew},
		"2": { jwtMasterOld, jwtMasterNew},
	})

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
			name: "Error for player old token",
			jwt: &Jwt{
				Secret:  jwtSecret,
				Storage: mockStorage,
			},
			args: args{
				request: JwtRequest{
					Token: jwtPlayerOld,
				},
			},
			want: JwtResponse{
				Payload: domain.Payload{},
			},
			wantErr: true,
		},
		{
			name: "Success for player new token",
			jwt: &Jwt{
				Secret:  jwtSecret,
				Storage: mockStorage,
			},
			args: args{
				request: JwtRequest{
					Token: jwtPlayerNew,
				},
			},
			want: JwtResponse{
				Payload: domain.Payload{
					Auth:    domain.RolePlayer,
					ModelId: 1,
					Exp:     expNew,
				},
			},
			wantErr: false,
		},
		{
			name: "Error for master old token",
			jwt: &Jwt{
				Secret:  jwtSecret,
				Storage: mockStorage,
			},
			args: args{
				request: JwtRequest{
					Token: jwtMasterOld,
				},
			},
			want: JwtResponse{
				Payload: domain.Payload{},
			},
			wantErr: true,
		},
		{
			name: "Success for master new token",
			jwt: &Jwt{
				Secret:  jwtSecret,
				Storage: mockStorage,
			},
			args: args{
				request: JwtRequest{
					Token: jwtMasterNew,
				},
			},
			want: JwtResponse{
				Payload: domain.Payload{
					Auth:    domain.RoleMaster,
					ModelId: 2,
					Exp:     expNew,
				},
			},
			wantErr: false,
		},
		{
			name: "Error secret",
			jwt: &Jwt{
				Secret:  "test2",
				Storage: mockStorage,
			},
			args: args{
				request: JwtRequest{
					Token: jwtPlayerOld,
				},
			},
			want: JwtResponse{
				Payload: domain.Payload{
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
