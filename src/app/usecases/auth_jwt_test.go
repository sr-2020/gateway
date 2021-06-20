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
	mockStorage := storage.NewMock()

	exp := time.Now().Add(5 * time.Minute).Unix()
	expPrev := time.Now().Add(1 * time.Minute).Unix()
	expNext := time.Now().Add(10 * time.Minute).Unix()

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
				Secret:  jwtSecret,
				Storage: mockStorage,
			},
			args: args{
				request: JwtRequest{
					Token: jwtToken(1, domain.RolePlayer, exp, jwtSecret),
				},
			},
			want: JwtResponse{
				Payload: domain.Payload{
					Auth:    domain.RolePlayer,
					ModelId: 1,
					Exp:     exp,
				},
			},
			wantErr: false,
		},
		{
			name: "Success for 1 next",
			jwt: &Jwt{
				Secret:  jwtSecret,
				Storage: mockStorage,
			},
			args: args{
				request: JwtRequest{
					Token: jwtToken(1, domain.RolePlayer, expNext, jwtSecret),
				},
			},
			want: JwtResponse{
				Payload: domain.Payload{
					Auth:    domain.RolePlayer,
					ModelId: 1,
					Exp:     expNext,
				},
			},
			wantErr: false,
		},
		{
			name: "Multi login error for 1 prev",
			jwt: &Jwt{
				Secret:  jwtSecret,
				Storage: mockStorage,
			},
			args: args{
				request: JwtRequest{
					Token: jwtToken(1, domain.RolePlayer, expPrev, jwtSecret),
				},
			},
			want: JwtResponse{
				Payload: domain.Payload{},
			},
			wantErr: true,
		},
		{
			name: "Success for 2",
			jwt: &Jwt{
				Secret:  jwtSecret,
				Storage: mockStorage,
			},
			args: args{
				request: JwtRequest{
					Token: jwtToken(2, domain.RoleMaster, exp, jwtSecret),
				},
			},
			want: JwtResponse{
				Payload: domain.Payload{
					Auth:    domain.RoleMaster,
					ModelId: 2,
					Exp:     exp,
				},
			},
			wantErr: false,
		},
		{
			name: "Success for 2 next",
			jwt: &Jwt{
				Secret:  jwtSecret,
				Storage: mockStorage,
			},
			args: args{
				request: JwtRequest{
					Token: jwtToken(2, domain.RoleMaster, expNext, jwtSecret),
				},
			},
			want: JwtResponse{
				Payload: domain.Payload{
					Auth:    domain.RoleMaster,
					ModelId: 2,
					Exp:     expNext,
				},
			},
			wantErr: false,
		},
		{
			name: "Success for 2 prev",
			jwt: &Jwt{
				Secret:  jwtSecret,
				Storage: mockStorage,
			},
			args: args{
				request: JwtRequest{
					Token: jwtToken(2, domain.RoleMaster, expPrev, jwtSecret),
				},
			},
			want: JwtResponse{
				Payload: domain.Payload{
					Auth:    domain.RoleMaster,
					ModelId: 2,
					Exp:     expPrev,
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
					Token: jwtToken(1, domain.RolePlayer, exp, jwtSecret),
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
