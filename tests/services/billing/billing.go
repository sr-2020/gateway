package billing

import (
	"encoding/json"
	"fmt"
	"github.com/sr-2020/gateway/tests/domain"
	"io/ioutil"
	"net/http"
)

type ServiceImpl struct {
	host string
	token string
}

func NewServiceImpl(host, token string) *ServiceImpl {
	return &ServiceImpl{host, token}
}

func (a *ServiceImpl) Check() bool {
	resp, _ := http.Get(fmt.Sprintf("%s/billing/test/testid", a.host))
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body) == "0"
}

func (a *ServiceImpl) Balance() (domain.BalanceResponse, error) {
	var balanceResponse domain.BalanceResponse

	client := http.Client{}
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/billing/sin", a.host),
		nil)
	if err != nil {
		return balanceResponse, err
	}

	req.Header.Set("Authorization", "Bearer " + a.token)

	resp, err := client.Do(req)
	if err != nil {
		return balanceResponse, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return balanceResponse, err
	}

	if err := json.Unmarshal(body, &balanceResponse); err != nil {
		return balanceResponse, err
	}

	return balanceResponse, nil
}
