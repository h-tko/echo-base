package controllers

// controller規定構造体
type Base struct {
	MetaTitle       string
	MetaDescription string
	MetaKeywords    string
	MetaH1          string
	MetaRobots      string

	response map[string]interface{}
}

// 事前処理
//
// param: c echo.Context
func (b *Base) BeforeFilter(c echo.Context) {
	b.clearAllResponse()
}

// 画面へのレスポンスパラメータ設定
//
// param: key string パラメータキー
// param: val interface{} パラメータ
func (b *Base) SetResponse(key string, val interface{}) {
	b.initResponse()

	b.response[key] = val
}

// htmlレンダリング
//
// param: c echo.Context
// param: status int HTTPレスポンスステータス
// param: oFile string レンダリング対象htmlファイル
// return: error
func (b *Base) Render(c echo.Context, status int, oFile string) error {
	b.initResponse()

	b.setMeta()

	return c.Render(status, oFile, b.response)
}

// JSONレスポンス
//
// param: c echo.Context
// param: status int HTTPレスポンスステータス
// return: error
func (b *Base) JSON(c echo.Context, status int) error {
	b.initResponse()

	b.setMeta()

	return c.JSON(status, b.response)
}

// Metaタグ設定
func (b *Base) setMeta() {
	b.response["mt"] = b.MetaTitle
	b.response["md"] = b.MetaDescription
	b.response["mk"] = b.MetaKeywords
	b.response["mh1"] = b.MetaH1
	b.response["mr"] = b.MetaRobots
}

// 全responseデータ削除
func (b *Base) clearAllResponse() {
	for key := range this.response {
		delete(this.response, key)
	}
}

// レスポンス用のインスタンスを初期化
func (b *Base) initResponse() {
	if b.response == nil {
		b.response = make(map[string]interface{})
	}
}
