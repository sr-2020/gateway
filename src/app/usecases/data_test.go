package usecases

import (
	"github.com/sr-2020/gateway/app/domain"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"

	"github.com/sr-2020/gateway/app/adapters/services/position"
)

func TestData_Execute(t *testing.T) {
	//cfg := config.LoadConfig()

	//positionService := position.NewService(cfg.Services["position"])
	positionService := new(position.MockService)

	type fields struct {
		Position position.ServiceInterface
	}
	type args struct {
		request DataRequest
	}
	tests := []struct {
		name    string
		mock    *mock.Call
		fields  fields
		args    args
		want    DataResponse
		wantErr bool
	}{
		{
			name: "Success with zero mana",
			mock: positionService.On("Location", 1).
				Return(domain.Location{
					Id:        1,
					ManaLevel: 0,
				}, nil),
			fields: fields{
				Position: positionService,
			},
			args: args{
				request: DataRequest{
					Id: 1,
					Scopes: []string{"position"},
				},
			},
			want: DataResponse{
				Data: map[string]interface{}{
					"position": domain.Location{
						Id:        1,
						ManaLevel: 0,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Success with non zero mana",
			mock: positionService.On("Location", 2).
				Return(domain.Location{
					Id:        2,
					ManaLevel: 100,
				}, nil),
			fields: fields{
				Position: positionService,
			},
			args: args{
				request: DataRequest{
					Id: 2,
					Scopes: []string{"position"},
				},
			},
			want: DataResponse{
				Data: map[string]interface{}{
					"position": domain.Location{
						Id:        2,
						ManaLevel: 100,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Not found",
			mock: positionService.On("Location", 404).
				Return(domain.Location{
					Id:        0,
					ManaLevel: 0,
				}, domain.ErrNotFound),
			fields: fields{
				Position: positionService,
			},
			args: args{
				request: DataRequest{
					Id: 404,
					Scopes: []string{"position"},
				},
			},
			want: DataResponse{
				Data: map[string]interface{}{
					"position": domain.Location{
						Id:        0,
						ManaLevel: 0,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Data{
				Position: tt.fields.Position,
			}
			got, err := d.Execute(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Data.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
