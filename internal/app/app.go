package app

import (
	"strings"
	"time"

	"github.com/booomch/demo-golang/internal/appconfig"
	"github.com/booomch/demo-golang/internal/controller"
	sconfiguration "github.com/booomch/demo-golang/internal/entities/_shared/configuration"
	userctr "github.com/booomch/demo-golang/internal/entities/user/controller"
	"github.com/booomch/demo-golang/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Setup(app *fiber.App, ctr *controller.Ctr) {
	p := middleware.MustParseClaims(ctr.TokenSvc())
	l := middleware.Limiter(&middleware.LimiterConfig{Max: 5, Expiration: time.Second})
	// Global middleware
	app.Use(recover.New())
	if appconfig.Config.Logger {
		app.Use(logger.New())
	}
	app.Use(cors.New())

	apiv1 := app.Group("/api/v1", p)
	usecaseEnpointsMap := make(map[interface{}][]sconfiguration.EndpointConfiguration)

	usecaseEnpointsMap[ctr.User] = userctr.ComposeEndpoints()

	for uc, endpoints := range usecaseEnpointsMap {
		for _, v := range endpoints {
			apiv1.Add(strings.ToUpper(v.Method), v.Path, handlerWrapper(*ctr, v.Handler, v.RequestDecoder, v.ResponseEncoder, uc))
		}
	}

}
