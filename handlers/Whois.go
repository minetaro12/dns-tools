package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/likexian/whois"
)

type whoisRequest struct {
	Domain string `json:"domain"`
}

func Whois(c *gin.Context) {
	// 応答をキャッシュしないように
	c.Header("cache-control", "public, max-age=0, must-revalidate")

	var reqBody whoisRequest

	// JSONのパースに失敗した場合はエラー
	if c.BindJSON(&reqBody) != nil {
		c.Writer.WriteString("Invalid Request")
		return
	}

	// whoisコマンドの実行
	whoisResult, err := whois.Whois(reqBody.Domain)

	// whoisコマンドの実行に失敗した場合はエラー
	if err != nil {
		c.Status(500)
		c.Writer.WriteString(err.Error())
		return
	}

	// レスポンス
	c.Writer.WriteString(whoisResult)
}
