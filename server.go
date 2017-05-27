package main

import (
	"fmt"
	"github.com/h-tko/echo-base/libraries"
	"github.com/joho/godotenv"
	"os"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.EecuteTemplate(w, name, data)
}

func envLoad() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	if os.Getenv("APP_ENV") == "" {
		os.Setenv("APP_ENV", "development")
	}

	return nil
}

func main() {
	// .env読み込み
	envLoad()

	conf, err := libraries.GetConfig()

	if err != nil {
		panic(err)
	}

	println(conf.Get("database.host"))
}