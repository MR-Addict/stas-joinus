package controllers

import (
	"server/configs"
	"server/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *fiber.Ctx) error {
	var user models.LoginCrential

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(models.Response{Success: false, Message: "表单内容不合法"})
	}

	if user.Username != configs.Env.ADMIN_USERNAME || user.Password != configs.Env.ADMIN_PASSWORD {
		return c.Status(400).JSON(models.Response{Success: false, Message: "用户名或密码错误"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"exp": time.Now().Add(time.Hour * 24 * 30).Unix()})
	s, err := token.SignedString([]byte(configs.Env.JWT_SECRET))
	if err != nil {
		return c.Status(500).JSON(models.Response{Success: false, Message: "无法生成token"})
	}

	return c.Status(200).JSON(models.Response{Success: true, Message: "登录成功", Token: s})
}
