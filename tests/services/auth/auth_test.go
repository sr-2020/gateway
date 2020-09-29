package auth

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/sr-2020/gateway/tests/config"
	"github.com/sr-2020/gateway/tests/domain"
	"net/http"
	"testing"
)

func TestCheck(t *testing.T) {
	convey.Convey("Go to check auth service", t, func() {
		cfg := config.LoadConfig()
		authService := NewServiceImpl(cfg.Host + "/api/v1")

		convey.Convey("Check response", func() {
			convey.So(authService.Check(), convey.ShouldEqual, true)
		})
	})
}

func TestLogin(t *testing.T) {
	cfg := config.LoadConfig()
	authService := NewServiceImpl(cfg.Host + "/api/v1")

	convey.Convey("Try to login with empty creds", t, func() {
		token, statusCode, err := authService.Auth(map[string]string{})

		convey.So(err, convey.ShouldBeNil)
		convey.So(statusCode, convey.ShouldEqual, http.StatusBadRequest)
		convey.So(token, convey.ShouldResemble, domain.Token{})
	})

	convey.Convey("Try to login for not exists account", t, func() {
		token, statusCode, err := authService.Auth(map[string]string{
			"login": "auth-wrong@test.com",
			"password": "1234",
		})

		convey.So(err, convey.ShouldBeNil)
		convey.So(statusCode, convey.ShouldEqual, http.StatusUnauthorized)
		convey.So(token, convey.ShouldResemble, domain.Token{})
	})

	convey.Convey("Try to login with wrong creds", t, func() {
		token, statusCode, err := authService.Auth(map[string]string{
			"login": cfg.Login,
			"password": "wrong-pass",
		})

		convey.So(err, convey.ShouldBeNil)
		convey.So(statusCode, convey.ShouldEqual, http.StatusUnauthorized)
		convey.So(token, convey.ShouldResemble, domain.Token{})
	})

	convey.Convey("Legacy login with valid creds with email", t, func() {
		token, statusCode, err := authService.Auth(map[string]string{
			"email":    cfg.Login,
			"password": cfg.Password,
		})

		convey.So(err, convey.ShouldBeNil)
		convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

		convey.So(token.Id, convey.ShouldEqual, cfg.ModelId)
		convey.So(token.ApiKey, convey.ShouldNotEqual, "")
	})

	convey.Convey("Login with valid creds", t, func() {
		token, statusCode, err := authService.Auth(map[string]string{
			"login": cfg.Login,
			"password": cfg.Password,
		})

		convey.So(err, convey.ShouldBeNil)
		convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

		convey.So(token.Id, convey.ShouldEqual, cfg.ModelId)
		convey.So(token.ApiKey, convey.ShouldNotEqual, "")

		modelId, err := authService.ModelId(token)
		convey.So(err, convey.ShouldBeNil)
		convey.So(modelId, convey.ShouldEqual, cfg.ModelId)

		oldToken := token
		convey.Convey("One more time login", func() {
			token, statusCode, err := authService.Auth(map[string]string{
				"login": cfg.Login,
				"password": cfg.Password,
			})

			convey.So(err, convey.ShouldBeNil)
			convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

			convey.So(token.Id, convey.ShouldEqual, cfg.ModelId)
			convey.So(token.ApiKey, convey.ShouldNotEqual, "")

			modelId, err := authService.ModelId(token)
			convey.So(err, convey.ShouldBeNil)
			convey.So(modelId, convey.ShouldEqual, cfg.ModelId)

			convey.Convey("Check auth with preview token", func() {
				modelId, err := authService.ModelId(oldToken)
				convey.So(err, convey.ShouldBeNil)
				convey.So(modelId, convey.ShouldEqual, cfg.ModelId)
			})
		})
	})
}
