package api

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"log"
	"net/http"
	"testing"
)

func TestAuthUsers(t *testing.T) {
	convey.Convey("Read users list", t, func() {

		client := http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/auth/users", cfg.Host), nil)
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
