package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/likexian/whois"
	"github.com/minetaro12/dns-tools/internal/models"
)

func PostWhois(c *fiber.Ctx) error {
	// 応答をキャッシュしない
	c.Set("cache-control", "public, max-age=0, must-revalidate")

	req := new(models.Whios)

	// JSONのパースに失敗した場合はエラー
	if err := c.BodyParser(req); err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(fiber.Map{
			"message": "Invalid JSON",
		})
		return err
	}

	// whoisの実行
	result, err := whois.Whois(req.Domain)

	// whoisの実行に失敗した場合はエラー
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.SendString(err.Error())
		return err
	}

	c.WriteString(result)
	return nil
}
