package GoBeer

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFormatter(t *testing.T) {
	Convey("Formatter", t, func() {
		time, _ := time.Parse("2006-01-02 15:04:05", "2015-03-04 23:29:23")
		formatedTime := Format(time, "%Y年%m月%d日 %H時%M分%S秒")
		So(formatedTime, ShouldEqual, "2015年03月04日 23時29分23秒")
	})
}
