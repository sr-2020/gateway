package tests

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/smartystreets/goconvey/convey"
	"github.com/sr-2020/gateway/tests/config"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

var (
	cfg config.Config
)

func jwtToken(id int, secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub":         "37445",
		"auth":        "ROLE_PLAYER",
		"modelId":     id,
		"characterId": 64,
		"exp":         time.Now().Add(1 * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}

	return tokenString
}

func TestMain(m *testing.M) {
	cfg = config.LoadConfig()

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestSuccess(t *testing.T) {
	convey.Convey("Go to check auth service", t, func() {

		client := http.Client{}
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/debug", cfg.Host), strings.NewReader(`{"test":"name"}`))
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwtToken(1, cfg.JwtSecret)))

		resp, _ := client.Do(req)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		var data map[string]map[string]interface{}
		if err := json.Unmarshal(body, &data); err != nil {
			return
		}

		convey.Convey("Check response", func() {
			convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)
			convey.So(data["headers"]["X-User-Id"], convey.ShouldResemble, []interface{}{"1"})

			convey.So(data["body"]["test"], convey.ShouldEqual, "name")
			convey.So(data["body"]["user_id"], convey.ShouldEqual, 1)
		})
	})
}

func TestEmptyJsonBody(t *testing.T) {
	convey.Convey("Go to check auth service", t, func() {

		client := http.Client{}
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/debug", cfg.Host), strings.NewReader(`{}`))
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwtToken(1, cfg.JwtSecret)))

		resp, _ := client.Do(req)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		var data map[string]map[string]interface{}
		if err := json.Unmarshal(body, &data); err != nil {
			return
		}

		convey.Convey("Check response", func() {
			convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)
			convey.So(data["headers"]["X-User-Id"], convey.ShouldResemble, []interface{}{"1"})

			convey.So(data["body"]["user_id"], convey.ShouldEqual, 1)
		})
	})
}

func TestNilBody(t *testing.T) {
	convey.Convey("Go to check auth service", t, func() {

		client := http.Client{}
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/debug", cfg.Host), nil)
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwtToken(1, cfg.JwtSecret)))

		resp, _ := client.Do(req)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		var data map[string]map[string]interface{}
		if err := json.Unmarshal(body, &data); err != nil {
			return
		}

		convey.Convey("Check response", func() {
			convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)
			convey.So(data["headers"]["X-User-Id"], convey.ShouldResemble, []interface{}{"1"})

			convey.So(data["body"]["user_id"], convey.ShouldEqual, nil)
		})
	})
}

func TestEmptyBody(t *testing.T) {
	convey.Convey("Go to check auth service", t, func() {

		client := http.Client{}
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/debug", cfg.Host), strings.NewReader(``))
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwtToken(1, cfg.JwtSecret)))

		resp, _ := client.Do(req)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		var data map[string]map[string]interface{}
		if err := json.Unmarshal(body, &data); err != nil {
			return
		}

		convey.Convey("Check response", func() {
			convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)
			convey.So(data["headers"]["X-User-Id"], convey.ShouldResemble, []interface{}{"1"})

			convey.So(data["body"]["user_id"], convey.ShouldEqual, nil)
		})
	})
}

func TestWrongJsonBody(t *testing.T) {
	convey.Convey("Go to check auth service", t, func() {

		client := http.Client{}
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/debug", cfg.Host), strings.NewReader(`{--`))
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwtToken(1, cfg.JwtSecret)))

		resp, _ := client.Do(req)
		defer resp.Body.Close()

		convey.Convey("Check response", func() {
			convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusBadRequest)
		})
	})
}
