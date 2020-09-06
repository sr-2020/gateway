package models_manager

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
	resp, _ := http.Get(fmt.Sprintf("%s/models-manager/ping", a.host))
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}
