package api

import (
	"encoding/json"
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/sr-2020/gateway/domain"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestModelsManagerPing(t *testing.T) {
	convey.Convey("Read models manager ping", t, func() {

		client := http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/models-manager/ping", cfg.Host), nil)
		if err != nil {
			log.Fatal(err)
			return
		}

		resp, _ := client.Do(req)
		defer resp.Body.Close()

		convey.Convey("Check OK status", func() {
			convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)
		})
	})
}

func TestModelsManagerCharacterModel(t *testing.T) {
	convey.Convey("Login with email and password success", t, func() {

		client := http.Client{}
		req, err := http.NewRequest("POST",
			fmt.Sprintf("%s/api/v1/auth/login", cfg.Host),
			strings.NewReader(fmt.Sprintf(`{"email":"%s","password":"%s"}`, login, password)))
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", "application/json")

		resp, _ := client.Do(req)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			return
		}

		var auth domain.Auth
		_ = json.Unmarshal(body, &auth)

		convey.Convey("Check response by model manager", func() {
			convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)

			client := http.Client{}
			req, err := http.NewRequest("GET",
				fmt.Sprintf("%s/api/v1/models-manager/character/model", cfg.Host),
				nil)
			if err != nil {
				return
			}

			req.Header.Set("Authorization", "Bearer " + auth.ApiKey)

			resp, _ := client.Do(req)
			defer resp.Body.Close()

			convey.Convey("Check OK status", func() {
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)
			})
		})
	})
}
