package libraries

import (
	"github.com/h-tko/echo-base/libraries"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetConfig(t *testing.T) {
	Convey("GetConfigのテスト", t, func() {
		Convey("正常系", func() {
			_, err := libraries.GetConfig()
			So(err, ShouldBeNil)
		})
	})
}
