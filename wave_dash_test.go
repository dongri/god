package god

import (
	"testing"
)

func TestReplaceWaveDash(t *testing.T) {
	result := ReplaceWaveDash("test῀test∼テスト～测试〜테스트")
	if result != "test~test~テスト~测试~테스트" {
		t.Fatal("Error!")
	}
}
