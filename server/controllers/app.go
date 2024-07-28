package controllers

import (
	"server/configs"
	"server/models"

	"github.com/gofiber/fiber/v2"
)

func AppConfig(c *fiber.Ctx) error {
	return c.Status(200).JSON(models.Response{Success: true, Message: "ok", Data: configs.Config.App})
}
