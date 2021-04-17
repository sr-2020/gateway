package billing

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/sr-2020/gateway/tests/config"
	"github.com/sr-2020/gateway/tests/services/auth"
	"testing"
)

func TestCheck(t *testing.T) {
	convey.Convey("Go to check billing service", t, func() {
		cfg := config.LoadConfig()
		billingService := NewServiceImpl(cfg.Host + "/api/v1", "")

		convey.Convey("Check response", func() {
			convey.So(billingService.Check(), convey.ShouldEqual, true)
		})
	})
}

func TestBalance(t *testing.T) {
	cfg := config.LoadConfig()
	authService := auth.NewServiceImpl(cfg.Host + "/api/v1")

	convey.Convey("Try to check balance for unauthorized user", t, func() {
		billingService := NewServiceImpl(cfg.Host+"/api/v1", "")

		balanceResponse, err := billingService.Balance()

		convey.So(err, convey.ShouldEqual, nil)
		convey.So(balanceResponse.Data.ModelId, convey.ShouldEqual, 0)
	})

	convey.Convey("Login with valid creds", t, func() {
		token, err := authService.AuthTest()
		convey.So(err, convey.ShouldBeNil)

		convey.Convey("Check balance for auth user", func() {
			billingService := NewServiceImpl(cfg.Host+"/api/v1", token.ApiKey)

			balanceResponse, err := billingService.Balance()

			convey.So(err, convey.ShouldEqual, nil)
			convey.So(balanceResponse.Data.ModelId, convey.ShouldEqual, cfg.ModelId)
		})
	})
}
