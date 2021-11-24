package gomen

import "testing"

func TestCipherString(t *testing.T) {
	encryptedText, err := EncryptString("passwordpasswordpasswordpassword", "hoge")
	if err != nil {
		t.Fatal(err)
	}

	originalText, err := DecryptString("passwordpasswordpasswordpassword", encryptedText)
	if err != nil {
		t.Fatal(err)
	}

	if originalText != "hoge" {
		t.Errorf("got %v\nwant %v", originalText, "hoge")
	}

}
