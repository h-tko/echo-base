package main

import (
	"fmt"
	"github.com/h-tko/echo-base/libraries"
	"github.com/h-tko/echo-base/route"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

// envLoad ...
//
// .envファイルのロード
func envLoad() error {
	err := godotenv.Load()

	if os.Getenv("APP_ENV") == "" {
		os.Setenv("APP_ENV", "development")
	}

	return err
}

// newEcho ...
//
// echoインスタンスの生成と初期化
func newEcho() *echo.Echo {
	e := echo.New()

	// ミドルウェア登録
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.AddTrailingSlash())

	// 静的ファイルのパス指定
	e.Static("/static", "assets")

	// template登録
	t := newTemplate()

	e.Renderer = t

	// routing設定
	route.Route(e)

	return e
}

// main ...
//
// エントリポイント
func main() {
	// .env読み込み
	envLoad()

	// echoインスタンス初期化
	e := newEcho()

	// configファイル読み出し
	conf, err := libraries.GetConfig()

	if err != nil {
		panic(err)
	}

	port := conf.GetString("application.port")

	// サーバー起動
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
