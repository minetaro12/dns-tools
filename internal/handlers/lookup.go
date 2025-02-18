package handlers

import (
	"net/http"

	"github.com/minetaro12/dns-tools/internal/lib"
	"github.com/minetaro12/dns-tools/internal/models"

	"github.com/gofiber/fiber/v2"
)

func PostLookup(c *fiber.Ctx) error {
	// 応答をキャッシュしないように
	c.Set("cache-control", "public, max-age=0, must-revalidate")

	req := new(models.Lookup)

	// JSONのパースに失敗した場合はエラー
	if err := c.BodyParser(req); err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(fiber.Map{
			"message": "Invalid JSON",
		})
		return err
	}

	// DNSサーバーのデフォルト設定
	if req.Dns == "" {
		req.Dns = "8.8.8.8"
	}

	// レコードタイプのデフォルト設定
	if req.Type == "" {
		req.Type = "A"
	}

	// lookupの実行
	result, err := lib.Resolve(*req)

	// lookupの実行に失敗した場合はエラー
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.SendString(err.Error())
		return err
	}

	c.WriteString(result)
	return nil
}
