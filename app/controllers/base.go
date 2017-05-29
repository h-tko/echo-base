package controllers

import (
	"github.com/labstack/echo"
)

// Base ...
//
// controller規定構造体.
//
// public functions:
//
//   BeforeFilter(c echo.Context)
//
//   SetResponse(key string, val interface{})
//
//   Render(c echo.Context, status int, oFile string) error
//
//   JSON(c echo.Context, status int) error
type Base struct {
	MetaTitle       string
	MetaDescription string
	MetaKeywords    string
	MetaH1          string
	MetaRobots      string

	response map[string]interface{}
}

// BeforeFilter ...
//
// 事前処理
func (b *Base) BeforeFilter(c echo.Context) {
	b.clearAllResponse()
}

// SetResponse ...
//
// 画面へのレスポンスパラメータ設定
func (b *Base) SetResponse(key string, val interface{}) {
	b.initResponse()

	b.response[key] = val
}

// Render ...
//
// htmlレンダリング
func (b *Base) Render(c echo.Context, status int, oFile string) error {
	b.initResponse()

	b.setMeta()

	return c.Render(status, oFile, b.response)
}

// JSON ...
//
// JSONレスポンス
func (b *Base) JSON(c echo.Context, status int) error {
	b.initResponse()

	b.setMeta()

	return c.JSON(status, b.response)
}

// setMeta ...
//
// Metaタグ設定
func (b *Base) setMeta() {
	b.response["mt"] = b.MetaTitle
	b.response["md"] = b.MetaDescription
	b.response["mk"] = b.MetaKeywords
	b.response["mh1"] = b.MetaH1
	b.response["mr"] = b.MetaRobots
}

// clearAllResponse ...
//
// 全responseデータ削除
func (b *Base) clearAllResponse() {
	for key := range b.response {
		delete(b.response, key)
	}
}

// initResponse ...
//
// レスポンス用のインスタンスを初期化
func (b *Base) initResponse() {
	if b.response == nil {
		b.response = make(map[string]interface{})
	}
}
