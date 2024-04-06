package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/likexian/whois"
)

type whoisRequest struct {
	Domain string `json:"domain"`
}

func Whois(c *fiber.Ctx) error {
	// 応答をキャッシュしない
	c.Set("cache-control", "public, max-age=0, must-revalidate")

	reqBody := new(whoisRequest)

	// JSONのパースに失敗した場合はエラー
	if err := c.BodyParser(reqBody); err != nil {
		c.Status(http.StatusBadRequest)
		c.SendString("Invalid Request")
		return err
	}

	// whoisの実行
	result, err := whois.Whois(reqBody.Domain)

	// whoisの実行に失敗した場合はエラー
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.SendString(err.Error())
		return err
	}

	c.WriteString(result)
	return nil
}
