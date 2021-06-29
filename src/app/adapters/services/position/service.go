package position

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sr-2020/gateway/app/adapters/config"
	"github.com/sr-2020/gateway/app/adapters/storage"
	"github.com/sr-2020/gateway/app/domain"
	"github.com/valyala/fasthttp"
	"strconv"
)

type Service struct {
	Config config.Service
	Client fasthttp.Client
	Store  storage.Storage
}

func NewService(config config.Service, store storage.Storage) *Service {

	client := fasthttp.Client{
		ReadTimeout: config.Timeout,
		WriteTimeout: config.Timeout,
	}

	return &Service{config, client, store}
}

func (s *Service) Location(id int) (domain.Location, error) {
	var location domain.Location

	positionCachePrefix := "position::"
	cacheKey := positionCachePrefix + strconv.Itoa(id)
	if s.Config.Cache != 0 {
		cacheLocation, err := s.Store.ReadCache(cacheKey)
		if err == nil {
			if err := json.Unmarshal([]byte(cacheLocation), &location); err == nil {
				return location, nil
			}
		}
	}

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetHost(s.Config.Host)
	req.SetRequestURI(fmt.Sprintf("%s/api/v1/users/%d", s.Config.Host, id))
	req.Header.Add("X-User-Id", "1")

	if err := s.Client.Do(req, resp); err != nil {
		return location, err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		if resp.StatusCode() == fasthttp.StatusNotFound {
			return location, domain.ErrNotFound
		}

		return location, errors.New(fmt.Sprintf("Unexpected status code: %d. Expecting %d", resp.StatusCode(), fasthttp.StatusOK))
	}

	var positionUser domain.PositionUser
	jsonErr := json.Unmarshal(resp.Body(), &positionUser)
	if jsonErr != nil {
		return location, jsonErr
	}

	if positionUser.Location != nil && positionUser.Location.Id != 0 {
		location.Id = positionUser.Location.Id
		if v, ok := positionUser.Location.Options["manaLevel"]; ok {
			if v, ok := v.(float64); ok {
				location.ManaLevel = int(v)
			}
		}
	}

	if s.Config.Cache != 0 {
		cacheValue, err := json.Marshal(location)
		if err != nil {
			// do nothing
		}

		if err := s.Store.WriteCache(cacheKey, cacheValue, s.Config.Cache); err != nil {
			// do nothing
		}
	}

	return location, nil
}



