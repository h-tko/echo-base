package main

import (
	"fmt"
	"github.com/h-tko/echo-base/libraries"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"html/template"
	"io"
	"os"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func envLoad() error {
	err := godotenv.Load()

	if os.Getenv("APP_ENV") == "" {
		os.Setenv("APP_ENV", "development")
	}

	return err
}

func main() {
	// .env読み込み
	envLoad()

	conf, err := libraries.GetConfig()

	if err != nil {
		panic(err)
	}

}
