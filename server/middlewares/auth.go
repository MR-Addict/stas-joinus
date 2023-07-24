package middlewares

import (
	"server/configs"
	"server/models"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

var Auth = jwtware.New(jwtware.Config{
	SigningKey: jwtware.SigningKey{Key: []byte(configs.Env.JWT_SECRET)},
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.Status(401).JSON(models.Response{Success: false, Message: "无效的token"})
	},
})
