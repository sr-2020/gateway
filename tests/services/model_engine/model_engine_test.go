package model_engine

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/sr-2020/gateway/tests/config"
	"testing"
)

func TestCheck(t *testing.T) {
	convey.Convey("Go to check model-engine service", t, func() {
		cfg := config.LoadConfig()
		modelEngineService := NewServiceImpl(cfg.Host + "/api/v1")

		convey.Convey("Check response", func() {
			convey.So(modelEngineService.Check(), convey.ShouldEqual, true)
		})
	})
}
