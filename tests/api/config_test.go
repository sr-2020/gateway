package api

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"net/http"
	"strings"
	"testing"
)

func TestConfigSuccess(t *testing.T) {
	convey.Convey("Set config test with key=value", t, func() {

		client := http.Client{}
		req, err := http.NewRequest("POST",
			fmt.Sprintf("%s/api/v1/config/test", cfg.Host),
			strings.NewReader(`{"key":"value"}`))
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", "application/json")

		resp, _ := client.Do(req)
		defer resp.Body.Close()

		convey.Convey("Check setting response", func() {
			convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)

			convey.Convey("Read config test", func() {
				convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)

				client := http.Client{}
				req, err := http.NewRequest("GET",
					fmt.Sprintf("%s/api/v1/config/test", cfg.Host),
					nil)
				if err != nil {
					return
				}
				req.Header.Set("Content-Type", "application/json")

				resp, _ := client.Do(req)
				defer resp.Body.Close()

				convey.Convey("Check reading response", func() {
					convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusOK)
				})

			})
		})
	})
}

func TestConfigFail(t *testing.T) {
	convey.Convey("Read not exists key", t, func() {

		client := http.Client{}
		req, err := http.NewRequest("GET",
			fmt.Sprintf("%s/api/v1/config/test404", cfg.Host),
			nil)
		if err != nil {
			return
		}

		resp, _ := client.Do(req)
		defer resp.Body.Close()

		convey.Convey("Check NotFound status", func() {
			convey.So(resp.StatusCode, convey.ShouldEqual, http.StatusNotFound)
		})
	})
}
