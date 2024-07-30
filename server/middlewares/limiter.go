package middlewares

import (
	"server/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

var Limitter = limiter.New(limiter.Config{
	Max:               2,
	Expiration:        10 * time.Second,
	LimiterMiddleware: limiter.SlidingWindow{},
	LimitReached: func(c *fiber.Ctx) error {
		return c.Status(429).JSON(models.Response{Success: false, Message: "你的请求太频繁啦，请稍后再试"})
	},
})
