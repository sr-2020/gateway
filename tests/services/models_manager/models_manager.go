package models_manager

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
	resp, _ := http.Get(fmt.Sprintf("%s/models-manager/ping", a.host))
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func (a *ServiceImpl) CharacterModel() (domain.CharacterModelResponse, error) {
	var characterModelResponse domain.CharacterModelResponse

	client := http.Client{}
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/models-manager/character/model", a.host),
		nil)
	if err != nil {
		return characterModelResponse, err
	}

	req.Header.Set("Authorization", "Bearer " + a.token)

	resp, err := client.Do(req)
	if err != nil {
		return characterModelResponse, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return characterModelResponse, domain.ErrUnauthorized
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return characterModelResponse, err
	}

	if err := json.Unmarshal(body, &characterModelResponse); err != nil {
		return characterModelResponse, err
	}

	return characterModelResponse, nil
}

func (a *ServiceImpl) SentEvent(event domain.Event) (domain.CharacterModelResponse, error) {
	var characterModelResponse domain.CharacterModelResponse

	if event.Data == nil {
		event.Data = make(map[string]interface{})
	}

	requestBody, err := json.Marshal(event)
	if err != nil {
		return characterModelResponse, err
	}
	dt := strings.NewReader(string(requestBody))

	client := http.Client{}
	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/models-manager/character/model", a.host),
		dt)
	if err != nil {
		return characterModelResponse, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + a.token)

	resp, err := client.Do(req)
	if err != nil {
		return characterModelResponse, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return characterModelResponse, domain.ErrUnauthorized
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return characterModelResponse, err
	}

	if err := json.Unmarshal(body, &characterModelResponse); err != nil {
		return characterModelResponse, err
	}

	return characterModelResponse, nil
}
