package gomen

import (
	"testing"
)

func TestDetectCardType(t *testing.T) {

	var cardType string
	var expected string

	// VISA ...
	cardType = DetectCardType("4111111111111")
	expected = "visa"
	if cardType != expected {
		t.Errorf("got %v\nwant %v", cardType, expected)
	}
	cardType = DetectCardType("4242424242424242")
	expected = "visa"
	if cardType != expected {
		t.Errorf("got %v\nwant %v", cardType, expected)
	}
	cardType = DetectCardType("4012888888881881")
	expected = "visa"
	if cardType != expected {
		t.Errorf("got %v\nwant %v", cardType, expected)
	}

	// MASTER ...
	cardType = DetectCardType("5555555555554444")
	expected = "master"
	if cardType != expected {
		t.Errorf("got %v\nwant %v", cardType, expected)
	}
	cardType = DetectCardType("5105105105105100")
	expected = "master"
	if cardType != expected {
		t.Errorf("got %v\nwant %v", cardType, expected)
	}
	cardType = DetectCardType("5431111111111111")
	expected = "master"
	if cardType != expected {
		t.Errorf("got %v\nwant %v", cardType, expected)
	}
	cardType = DetectCardType("5111111111111118")
	expected = "master"
	if cardType != expected {
		t.Errorf("got %v\nwant %v", cardType, expected)
	}

	// JCB ...
	cardType = DetectCardType("3530111333300000")
	expected = "jcb"
	if cardType != expected {
		t.Errorf("got %v\nwant %v", cardType, expected)
	}
	cardType = DetectCardType("3566002020360505")
	expected = "jcb"
	if cardType != expected {
		t.Errorf("got %v\nwant %v", cardType, expected)
	}

	// AMEX ...
	cardType = DetectCardType("378282246310005")
	expected = "amex"
	if cardType != expected {
		t.Errorf("got %v\nwant %v", cardType, expected)
	}
	cardType = DetectCardType("371449635398431")
	expected = "amex"
	if cardType != expected {
		t.Errorf("got %v\nwant %v", cardType, expected)
	}
	cardType = DetectCardType("341111111111111")
	expected = "amex"
	if cardType != expected {
		t.Errorf("got %v\nwant %v", cardType, expected)
	}
	cardType = DetectCardType("378734493671000")
	expected = "amex"
	if cardType != expected {
		t.Errorf("got %v\nwant %v", cardType, expected)
	}

	// DISCOVER ...
	cardType = DetectCardType("30569309025904")
	expected = "diners"
	if cardType != expected {
		t.Errorf("got %v\nwant %v", cardType, expected)
	}
	cardType = DetectCardType("38520000023237")
	expected = "diners"
	if cardType != expected {
		t.Errorf("got %v\nwant %v", cardType, expected)
	}

	// DISCOVER ...
	cardType = DetectCardType("6011111111111117")
	expected = "discover"
	if cardType != expected {
		t.Errorf("got %v\nwant %v", cardType, expected)
	}
	cardType = DetectCardType("6011000990139424")
	expected = "discover"
	if cardType != expected {
		t.Errorf("got %v\nwant %v", cardType, expected)
	}
	cardType = DetectCardType("6011601160116611")
	expected = "discover"
	if cardType != expected {
		t.Errorf("got %v\nwant %v", cardType, expected)
	}

}
