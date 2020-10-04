package models_manager

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/sr-2020/gateway/tests/config"
	"github.com/sr-2020/gateway/tests/domain"
	"github.com/sr-2020/gateway/tests/services/auth"
	"strconv"
	"testing"
)

func TestCheck(t *testing.T) {
	convey.Convey("Go to check models-manager service", t, func() {
		cfg := config.LoadConfig()
		modelsManagerService := NewServiceImpl(cfg.Host + "/api/v1", "")

		convey.Convey("Check response", func() {
			convey.So(modelsManagerService.Check(), convey.ShouldEqual, true)
		})
	})
}

func TestCharacterModel(t *testing.T) {
	cfg := config.LoadConfig()
	authService := auth.NewServiceImpl(cfg.Host + "/api/v1")

	convey.Convey("Login with valid creds", t, func() {
		token, err := authService.AuthTest()
		convey.So(err, convey.ShouldBeNil)

		convey.Convey("Read character model", func() {
			cfg := config.LoadConfig()
			modelsManagerService := NewServiceImpl(cfg.Host+"/api/v1", token.ApiKey)

			convey.Convey("Check response", func() {
				modelsManagerResponse, err := modelsManagerService.CharacterModel()

				convey.So(err, convey.ShouldEqual, nil)
				convey.So(modelsManagerResponse.BaseModel.ModelId, convey.ShouldEqual, strconv.Itoa(cfg.ModelId))
			})
		})
	})
}

func TestSentEventRevive(t *testing.T) {
	cfg := config.LoadConfig()
	authService := auth.NewServiceImpl(cfg.Host + "/api/v1")

	convey.Convey("Login with valid creds", t, func() {
		token, err := authService.AuthTest()
		convey.So(err, convey.ShouldBeNil)

		convey.Convey("Sent revive", func() {
			modelsManagerService := NewServiceImpl(cfg.Host+"/api/v1", token.ApiKey)

			event := domain.Event{
				EventType: "revive",
			}
			modelsManagerResponse, err := modelsManagerService.SentEvent(event)

			convey.So(err, convey.ShouldEqual, nil)
			convey.So(modelsManagerResponse.BaseModel.ModelId, convey.ShouldEqual, strconv.Itoa(cfg.ModelId))
		})
	})
}