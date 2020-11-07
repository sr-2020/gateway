package position

import (
	"encoding/json"
	"fmt"
	"github.com/sr-2020/gateway/tests/domain"
	"io/ioutil"
	"net/http"
	"strings"
)

type ServiceImpl struct {
	host string
	token string
}

func NewServiceImpl(host, token string) *ServiceImpl {
	return &ServiceImpl{host, token}
}

func (a *ServiceImpl) Check() bool {
	resp, _ := http.Get(fmt.Sprintf("%s/position/version", a.host))
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body) == "2010"
}

func (a *ServiceImpl) Locations() ([]domain.Location, error) {
	var locations []domain.Location

	client := http.Client{}
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/position/locations", a.host),
		nil)
	if err != nil {
		return locations, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return locations, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return locations, err
	}

	if err := json.Unmarshal(body, &locations); err != nil {
		return locations, err
	}

	return locations, nil
}

func (a *ServiceImpl) AddPosition(beacons []domain.Beacon) (domain.Position, error) {
	var position domain.Position

	reqBeacons := domain.Beacons{
		Beacons: beacons,
	}

	requestBody, err := json.Marshal(reqBeacons)
	if err != nil {
		return position, err
	}
	dt := strings.NewReader(string(requestBody))

	client := http.Client{}
	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/position/positions", a.host),
		dt)
	if err != nil {
		return position, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + a.token)

	resp, err := client.Do(req)
	if err != nil {
		return position, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return position, err
	}

	if err := json.Unmarshal(body, &position); err != nil {
		return position, err
	}

	return position, nil
}

func (a *ServiceImpl) ManaLevel() (domain.ManaLevel, error) {
	var manaLevel domain.ManaLevel

	client := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/position/manalevel", a.host), nil)
	if err != nil {
		return manaLevel, err
	}

	req.Header.Set("Authorization", "Bearer " + a.token)

	resp, err := client.Do(req)
	if err != nil {
		return manaLevel, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return manaLevel, domain.ErrUnauthorized
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return manaLevel, err
	}

	if err := json.Unmarshal(body, &manaLevel); err != nil {
		return manaLevel, err
	}

	return manaLevel, nil
}
