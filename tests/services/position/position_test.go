package position

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/sr-2020/gateway/tests/config"
	"github.com/sr-2020/gateway/tests/domain"
	"github.com/sr-2020/gateway/tests/services/auth"
	"testing"
)

func TestCheck(t *testing.T) {
	convey.Convey("Go to check position service", t, func() {
		cfg := config.LoadConfig()
		positionService := NewServiceImpl(cfg.Host + "/api/v1", "")

		convey.So(positionService.Check(), convey.ShouldEqual, true)
	})
}

func TestLocations(t *testing.T) {
	convey.Convey("Read list of locations", t, func() {
		cfg := config.LoadConfig()
		positionService := NewServiceImpl(cfg.Host+"/api/v1", "")

		_, err := positionService.Locations()

		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestAddPosition(t *testing.T) {
	cfg := config.LoadConfig()
	authService := auth.NewServiceImpl(cfg.Host + "/api/v1")

	convey.Convey("Login with valid creds", t, func() {
		token, err := authService.AuthTest()
		convey.So(err, convey.ShouldBeNil)

		convey.Convey("Add position for 3062 location", func() {
			positionService := NewServiceImpl(cfg.Host+"/api/v1", token.ApiKey)
			beacons := []domain.Beacon{
				{
					SSID: "",
					BSSID: "55:55:55:55:55:55",
					Level: -50,
				},
			}
			position, err := positionService.AddPosition(beacons)

			convey.So(err, convey.ShouldEqual, nil)
			convey.So(position.UserId, convey.ShouldEqual, cfg.ModelId)
			convey.So(position.LocationId, convey.ShouldEqual, 3062)

			convey.Convey("Add position for 3080 location", func() {
				beacons := []domain.Beacon{
					{
						SSID: "",
						BSSID: "22:22:22:22:22:22",
						Level: -50,
					},
				}
				position, err := positionService.AddPosition(beacons)

				convey.So(err, convey.ShouldEqual, nil)
				convey.So(position.UserId, convey.ShouldEqual, cfg.ModelId)
				convey.So(position.LocationId, convey.ShouldEqual, 3080)
			})
		})
	})
}
