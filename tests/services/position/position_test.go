package position

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/sr-2020/gateway/tests/config"
	"testing"
)

func TestCheck(t *testing.T) {
	convey.Convey("Go to check position service", t, func() {
		cfg := config.LoadConfig()
		positionService := NewServiceImpl(cfg.Host + "/api/v1")

		convey.Convey("Check response", func() {
			convey.So(positionService.Check(), convey.ShouldEqual, true)
		})
	})
}
