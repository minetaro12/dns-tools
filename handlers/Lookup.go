package handlers

import (
	"dns-tools/lib"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type lookupRequest struct {
	FQDN string `json:"fqdn"`
	DNS  string `json:"dns"`
	TYPE string `json:"type"`
}

func Lookup(c *fiber.Ctx) error {
	// 応答をキャッシュしないように
	c.Set("cache-control", "public, max-age=0, must-revalidate")

	reqBody := new(lookupRequest)

	// JSONのパースに失敗した場合はエラー
	if err := c.BodyParser(reqBody); err != nil {
		c.Status(http.StatusBadRequest)
		c.SendString("Invalid Request")
		return err
	}

	// lookupの実行
	result, err := execLookup(*reqBody)

	// lookupの実行に失敗した場合はエラー
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.SendString(err.Error())
		return err
	}

	c.WriteString(result)
	return nil
}

func execLookup(t lookupRequest) (string, error) {
	// 空ならば8.8.8.8を設定する
	if t.DNS == "" {
		t.DNS = "8.8.8.8"
	}

	// 空ならばAレコードに設定する
	if t.TYPE == "" {
		t.TYPE = "A"
	}

	r, err := lib.Resolve(t.FQDN, t.DNS, t.TYPE)
	if err != nil {
		return "", err
	}
	return r, nil
}
