package god

import (
	"testing"
)

func TestPackBeacon(t *testing.T) {
	major, minor := PackBeacon(1, 2, 8)
	expectedMajor, expectedMinor := 0, 258
	if major != expectedMajor || minor != expectedMinor {
		t.Errorf("got %v\nwant %v", major, expectedMajor)
		t.Errorf("got %v\nwant %v", minor, expectedMinor)
	}
}

func TestUnpackBeacon(t *testing.T) {
	major, minor := UnpackBeacon(0, 258, 8)
	expectedMajor, expectedMinor := 1, 2
	if major != expectedMajor || minor != expectedMinor {
		t.Errorf("got %v\nwant %v", major, expectedMajor)
		t.Errorf("got %v\nwant %v", minor, expectedMinor)
	}
}
