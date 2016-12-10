package utils

import (
	"testing"
	"time"
)

func TestDateFormat(t *testing.T) {
	time, _ := time.Parse("2006-01-02 15:04:05", "2015-03-04 23:29:23")
	formatedTime := DateFormat(time, "%Y年%m月%d日 %H時%M分%S秒")
	expected := "2015年03月04日 23時29分23秒"
	if formatedTime != expected {
		t.Errorf("got %v\nwant %v", formatedTime, expected)
	}
}
