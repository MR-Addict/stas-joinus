package middlewares

import (
	"server/configs"
	"server/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/golang-jwt/jwt/v5"
)

var Auth = keyauth.New(keyauth.Config{
	KeyLookup: "cookie:joinus_token",
	Validator: func(c *fiber.Ctx, key string) (bool, error) {
		token, err := jwt.Parse(key, func(token *jwt.Token) (interface{}, error) {
			return []byte(configs.Env.JWT_SECRET), nil
		})
		if err != nil {
			return false, keyauth.ErrMissingOrMalformedAPIKey
		}
		// check whether has token expired
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
			if expirationTime.Before(time.Now()) {
				return false, keyauth.ErrMissingOrMalformedAPIKey
			}
			return true, nil
		}
		return false, keyauth.ErrMissingOrMalformedAPIKey
	},
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.Status(401).JSON(models.Response{Success: false, Message: "无效的token"})
	},
})
