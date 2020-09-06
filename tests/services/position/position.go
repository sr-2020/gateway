package position

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type ServiceImpl struct {
	host string
}

func NewServiceImpl(host string) *ServiceImpl {
	return &ServiceImpl{host}
}

func (a *ServiceImpl) Check() bool {
	resp, _ := http.Get(fmt.Sprintf("%s/position/version", a.host))
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body) == "2010"
}
