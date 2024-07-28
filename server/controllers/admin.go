package controllers

import (
	"server/configs"
	"server/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func UserPing(c *fiber.Ctx) error {
	return c.Status(200).JSON(models.Response{Success: true, Message: "ok"})
}

func UserLogout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "joinus_token",
		Value:    "",
		Path:     "/",
		SameSite: "Lax",
		HTTPOnly: true,
		Expires:  time.Now().Add(-time.Hour),
	}

	c.Cookie(&cookie)

	return c.Status(200).JSON(models.Response{Success: true, Message: "退出成功"})
}

func UserLogin(c *fiber.Ctx) error {
	type LoginCrential struct {
		Password string `json:"password"`
	}

	var user LoginCrential

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(models.Response{Success: false, Message: "无效的请求体"})
	}

	if user.Password != configs.Config.AdminPassword {
		return c.Status(400).JSON(models.Response{Success: false, Message: "登录密码错误"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"exp": time.Now().Add(time.Hour * 24 * 30).Unix()})
	s, err := token.SignedString([]byte(configs.Config.AdminPassword))
	if err != nil {
		return c.Status(500).JSON(models.Response{Success: false, Message: "无法生成token"})
	}

	cookie := fiber.Cookie{
		Name:     "joinus_token",
		Value:    s,
		Path:     "/",
		SameSite: "Lax",
		HTTPOnly: true,
		MaxAge:   60 * 60 * 24 * 30,
	}

	c.Cookie(&cookie)

	return c.Status(200).JSON(models.Response{Success: true, Message: "登录成功", Data: s})
}
