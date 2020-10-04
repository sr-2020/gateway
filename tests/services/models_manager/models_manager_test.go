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

		convey.So(modelsManagerService.Check(), convey.ShouldEqual, true)
	})
}

func TestCharacterModel(t *testing.T) {
	cfg := config.LoadConfig()
	authService := auth.NewServiceImpl(cfg.Host + "/api/v1")

	convey.Convey("Try to read character model for unauthorized user", t, func() {
		modelsManagerService := NewServiceImpl(cfg.Host+"/api/v1", "")

		_, err := modelsManagerService.CharacterModel()

		convey.So(err, convey.ShouldEqual, domain.ErrUnauthorized)
	})

	convey.Convey("Login with valid creds", t, func() {
		token, err := authService.AuthTest()
		convey.So(err, convey.ShouldBeNil)

		convey.Convey("Read character model", func() {
			cfg := config.LoadConfig()
			modelsManagerService := NewServiceImpl(cfg.Host+"/api/v1", token.ApiKey)

			modelsManagerResponse, err := modelsManagerService.CharacterModel()

			convey.So(err, convey.ShouldEqual, nil)
			convey.So(modelsManagerResponse.BaseModel.ModelId, convey.ShouldEqual, strconv.Itoa(cfg.ModelId))
		})
	})
}

func TestSentEventRevive(t *testing.T) {
	cfg := config.LoadConfig()
	authService := auth.NewServiceImpl(cfg.Host + "/api/v1")

	convey.Convey("Try to sent revive for unauthorized user", t, func() {
		modelsManagerService := NewServiceImpl(cfg.Host+"/api/v1", "")

		event := domain.Event{
			EventType: "revive",
		}
		_, err := modelsManagerService.SentEvent(event)

		convey.So(err, convey.ShouldEqual, domain.ErrUnauthorized)
	})

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