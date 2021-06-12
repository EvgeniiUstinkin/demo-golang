package middleware

import (
	"net/http"
	"time"

	"github.com/booomch/demo-golang/pkg/codes"
	"github.com/booomch/demo-golang/pkg/httperr"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type LimiterConfig struct {
	Max        int
	Expiration time.Duration
}

func Limiter(config *LimiterConfig) fiber.Handler {
	return limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        config.Max,
		Expiration: config.Expiration,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return httperr.New(codes.Omit, http.StatusTooManyRequests, http.StatusText(http.StatusTooManyRequests)).Send(c)
		},
	})
}
