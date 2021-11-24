package gou

import (
	"testing"
)

func TestMinOf(t *testing.T) {
	min := MinOf(4, 6, 9, 7, 2, 1, 8)
	if min != 1 {
		t.Fatal("Error!")
	}
}

func TestMaxOf(t *testing.T) {
	max := MaxOf(4, 6, 9, 7, 2, 1, 8)
	if max != 9 {
		t.Fatal("Error!")
	}
}
