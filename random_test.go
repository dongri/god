package gou

import "testing"

func TestRandomStringDigits(t *testing.T) {
	random := new(Random)
	random.UseNumber()
	random.UseSmallLetter()
	random.UseCapitalLetter()
	r := random.Random(20)
	if len(r) != 20 {
		t.Fatal("Wrong length returned")
	}
}
