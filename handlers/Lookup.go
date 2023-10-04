package handlers

import (
	"dns-tools/lib"

	"github.com/gin-gonic/gin"
)

type lookupRequest struct {
	FQDN string `json:"fqdn"`
	DNS  string `json:"dns"`
	TYPE string `json:"type"`
}

func Lookup(c *gin.Context) {
	// 応答をキャッシュしないように
	c.Header("cache-control", "public, max-age=0, must-revalidate")

	var reqBody lookupRequest

	// JSONのパースに失敗した場合はエラー
	if c.BindJSON(&reqBody) != nil {
		c.Writer.WriteString("Invalid Request")
		return
	}

	// lookupの実行
	lookupResult, err := execLookup(reqBody)

	// lookupの実行に失敗した場合はエラー
	if err != nil {
		c.Status(500)
		c.Writer.WriteString(err.Error())
		return
	}

	// レスポンス
	c.Writer.WriteString(lookupResult)
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
