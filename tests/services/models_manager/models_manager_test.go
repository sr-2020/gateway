package models_manager

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/sr-2020/gateway/tests/config"
	"testing"
)

func TestCheck(t *testing.T) {
	convey.Convey("Go to check models-manager service", t, func() {
		cfg := config.LoadConfig()
		modelsManagerService := NewServiceImpl(cfg.Host + "/api/v1")

		convey.Convey("Check response", func() {
			convey.So(modelsManagerService.Check(), convey.ShouldEqual, true)
		})
	})
}
