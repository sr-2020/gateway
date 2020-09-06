package usecases

import (
	"github.com/sr-2020/gateway/app/adapters/services/position"
	"github.com/sr-2020/gateway/app/domain"
)

type DataInterface interface {
	Execute(DataRequest) (DataResponse, error)
}

type DataRequest struct {
	Id     int
	Scopes []string
}

type DataResponse struct {
	Data map[string]interface{}
}

type Data struct {
	Position position.ServiceInterface
}

func (d *Data) Execute(request DataRequest) (DataResponse, error) {
	var response DataResponse

	if len(request.Scopes) == 0 || request.Scopes[0] != "position" {
		return response, nil
	}

	response.Data = make(map[string]interface{})

	location, err := d.Position.Location(request.Id)
	if err != nil {
		if err != domain.ErrNotFound {
			return response, err
		}
	}

	response.Data["position"] = location

	return response, nil
}
