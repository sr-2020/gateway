package config

import (
	"errors"
	"fmt"
	"github.com/sr-2020/gateway/tests/domain"
	"io/ioutil"
	"net/http"
	"strings"
)

type ServiceImpl struct {
	host string
}

func NewServiceImpl(host string) *ServiceImpl {
	return &ServiceImpl{host}
}

func (a *ServiceImpl) Check() bool {
	resp, _ := http.Get(fmt.Sprintf("%s/config/test404", a.host))
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusNotFound
}

func (a *ServiceImpl) Read(key string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/config/%s", a.host, key))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode == http.StatusNotFound {
		return "", domain.ErrNotFound
	}

	return string(body), nil
}

func (a *ServiceImpl) Write(key, value string) error {
	dt := strings.NewReader(value)
	resp, err := http.Post(fmt.Sprintf("%s/config/%s", a.host, key), "application/json", dt)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return domain.ErrBadRequest
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Wrong status code %d", resp.StatusCode))
	}

	return nil
}
