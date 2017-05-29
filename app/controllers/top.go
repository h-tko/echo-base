package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

// Top ...
//
// TopController.
type Top struct {
	Base
}

// Index ...
//
// top画面表示処理.
func (t *Top) Index(c echo.Context) error {
	t.SetResponse("hoge", "moge")

	return t.JSON(c, http.StatusOK)
}
