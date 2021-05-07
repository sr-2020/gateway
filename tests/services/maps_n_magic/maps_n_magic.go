package maps_n_magic

import (
	"fmt"
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
	resp, _ := http.Get(fmt.Sprintf("%s/maps-n-magic/manifest.json", a.host))
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func (a *ServiceImpl) FileList() bool {
	resp, _ := http.Get(fmt.Sprintf("%s/maps-n-magic/fileList", a.host))
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}