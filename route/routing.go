package route

import (
	"github.com/h-tko/echo-base/app/controllers"
	"github.com/labstack/echo"
)

// Route ...
//
// routing設定
func Route(e *echo.Echo) {
	e.GET("/", (&controllers.Top{}).Index)
}
