package api

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"log"
	"net/http"
	"testing"
)

func TestModelEnginePing(t *testing.T) {
	convey.Convey("Read model engine ping", t, func() {

		client := http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/model-engine/ping", cfg.Host), nil)
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
