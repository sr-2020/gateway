package config

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/sr-2020/gateway/tests/config"
	"github.com/sr-2020/gateway/tests/domain"
	"testing"
)

func TestCheck(t *testing.T) {
	convey.Convey("Go to check config service", t, func() {
		cfg := config.LoadConfig()
		configService := NewServiceImpl(cfg.Host + "/api/v1")

		convey.Convey("Check response", func() {
			convey.So(configService.Check(), convey.ShouldEqual, true)
		})
	})
}

func TestWriteAndRead(t *testing.T) {
	convey.Convey("Write to config service testkey", t, func() {
		cfg := config.LoadConfig()
		configService := NewServiceImpl(cfg.Host + "/api/v1")

		jsonValue := `{"data":"value"}`
		convey.Convey("Check response for write", func() {
			convey.So(configService.Write("testkey", jsonValue), convey.ShouldEqual, nil)

			content, err := configService.Read("testkey")

			convey.So(content, convey.ShouldEqual, jsonValue)
			convey.So(err, convey.ShouldEqual, nil)
		})
	})
}

func TestWriteInvalidJson(t *testing.T) {
	convey.Convey("Try to Write invalid json value", t, func() {
		cfg := config.LoadConfig()
		configService := NewServiceImpl(cfg.Host + "/api/v1")

		jsonValue := `rawdata`
		convey.So(configService.Write("testkey", jsonValue), convey.ShouldEqual, domain.ErrBadRequest)
	})
}

func TestReadNotFound(t *testing.T) {
	convey.Convey("Read not found key in config service", t, func() {
		cfg := config.LoadConfig()
		configService := NewServiceImpl(cfg.Host + "/api/v1")

		content, err := configService.Read("testkey404")

		convey.So(content, convey.ShouldEqual, "")
		convey.So(err, convey.ShouldEqual, domain.ErrNotFound)
	})
}
