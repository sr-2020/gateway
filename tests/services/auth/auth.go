package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sr-2020/gateway/tests/config"
	"github.com/sr-2020/gateway/tests/domain"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type ServiceImpl struct {
	host string
}

func NewServiceImpl(host string) *ServiceImpl {
	return &ServiceImpl{host}
}

func (a *ServiceImpl) Check() bool {
	resp, _ := http.Get(fmt.Sprintf("%s/auth/login", a.host))
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func (a *ServiceImpl) Auth(data map[string]string) (domain.Token, int, error) {
	requestBody, _ := json.Marshal(data)

	dt := strings.NewReader(string(requestBody))
	resp, _ := http.Post(fmt.Sprintf("%s/auth/login", a.host), "application/json", dt)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	token := domain.Token{}
	if resp.StatusCode != http.StatusOK {
		return token, resp.StatusCode, errors.New("Status Code:" + strconv.Itoa(resp.StatusCode))
	}

	if err := json.Unmarshal(body, &token); err != nil {
		return token, resp.StatusCode, err
	}

	return token, resp.StatusCode, nil
}

func (a *ServiceImpl) AuthTest() (domain.Token, error) {
	cfg := config.LoadConfig()

	token, _, err := a.Auth(map[string]string{
		"login": cfg.Login,
		"password": cfg.Password,
	})

	return token, err
}

func (a *ServiceImpl) ModelId(token domain.Token) (int, error) {
	modelId := 0

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/billing/test/testid", a.host), nil)
	if err != nil {
		return modelId, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.ApiKey))

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return modelId, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return modelId, err
	}
	defer resp.Body.Close()

	modelId, err = strconv.Atoi(string(body))
	if err != nil {
		return modelId, err
	}

	return modelId, nil
}
