package billing

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/sr-2020/gateway/tests/config"
	"testing"
)

func TestCheck(t *testing.T) {
	convey.Convey("Go to check billing service", t, func() {
		cfg := config.LoadConfig()
		billingService := NewServiceImpl(cfg.Host + "/api/v1")

		convey.Convey("Check response", func() {
			convey.So(billingService.Check(), convey.ShouldEqual, true)
		})
	})
}
