package maps_n_magic

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/sr-2020/gateway/tests/config"
	"github.com/sr-2020/gateway/tests/services/auth"
	"testing"
)

func TestCheck(t *testing.T) {
	convey.Convey("Go to check maps-n-magic service", t, func() {
		cfg := config.LoadConfig()
		mapsMagicService := NewServiceImpl(cfg.Host + "/api/v1", "")

		convey.So(mapsMagicService.Check(), convey.ShouldEqual, true)
	})
}

func TestFileList(t *testing.T) {
	cfg := config.LoadConfig()
	authService := auth.NewServiceImpl(cfg.Host + "/api/v1")

	convey.Convey("Try to read file list for unauthorized user", t, func() {
		mapsMagicService := NewServiceImpl(cfg.Host+"/api/v1", "")

		convey.So(mapsMagicService.FileList(), convey.ShouldEqual, false)
	})

	convey.Convey("Login with valid creds", t, func() {
		token, err := authService.AuthTest()
		convey.So(err, convey.ShouldBeNil)

		convey.Convey("Try to read file list for not admin user", func() {
			mapsMagicService := NewServiceImpl(cfg.Host+"/api/v1", token.ApiKey)

			convey.So(mapsMagicService.FileList(), convey.ShouldEqual, false)
		})
	})
}