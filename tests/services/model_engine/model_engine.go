package model_engine

import (
	"fmt"
	"net/http"
)

type ServiceImpl struct {
	host string
}

func NewServiceImpl(host string) *ServiceImpl {
	return &ServiceImpl{host}
}

func (a *ServiceImpl) Check() bool {
	resp, _ := http.Get(fmt.Sprintf("%s/model-engine/ping", a.host))
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}
