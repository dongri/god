package goutils

import (
	"fmt"
	"testing"
)

func TestRandomStringDigits(t *testing.T) {
	random := new(Random)
	random.UseNumber()
	random.UseSmallLetter()
	random.UseCapitalLetter()
	r := random.Random(20)
	fmt.Println(r)
	if len(r) != 20 {
		t.Fatal("Wrong length returned")
	}
}
