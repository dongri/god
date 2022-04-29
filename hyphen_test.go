package god

import (
	"fmt"
	"testing"
)

func TestReplaceHyphen(t *testing.T) {
	result := ReplaceHyphen("test-testーテスト█测试▬테스트", HANKAKU_HYPHEN)
	if result != "test-test-テスト-测试-테스트" {
		t.Fatal("Error!")
	}
	result = ReplaceHyphen("test-testーテスト█测试▬테스트", ZENKAKU_HYPHEN)
	if result != "testーtestーテストー测试ー테스트" {
		fmt.Println(result)
		t.Fatal("Error!")
	}
}
