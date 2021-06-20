package app

import (
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sr-2020/gateway/app/adapters/config"
	"github.com/sr-2020/gateway/app/adapters/services/position"
	"github.com/sr-2020/gateway/app/adapters/storage"
	"github.com/sr-2020/gateway/app/handlers"
	"github.com/sr-2020/gateway/app/usecases"
	"strconv"
)

func DebugRoute(e *echo.Echo) {
	path := "/debug"

	e.GET(path, handlers.Debug)
	e.POST(path, handlers.Debug)
	e.PUT(path, handlers.Debug)
	e.PATCH(path, handlers.Debug)
	e.DELETE(path, handlers.Debug)
	e.HEAD(path, handlers.Debug)
	e.OPTIONS(path, handlers.Debug)
}

func Start(cfg config.Config) error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	redisStore := storage.NewRedis(redis.NewClient(&cfg.Redis))

	authHandler := handlers.Auth{
		UseCase: &usecases.Jwt{
			Secret: cfg.JwtSecret,
			Storage: redisStore,
		},
		UseCaseData: &usecases.Data{
			Position: position.NewService(cfg.Services["position"]),
		},
	}

	// Routes
	e.GET("/auth", authHandler.Handler)
	DebugRoute(e)

	// Start server
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(cfg.Port)))

	return nil
}
