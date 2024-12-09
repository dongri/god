package god

import (
	"strings"
	"unicode"
)

const alphabetLength = 26
const digitLength = 10

type Cipher interface {
	Encrypt(text string) string
	Decrypt(text string) string
}

type CaesarCipher struct {
	shift int
}

func NewCaesarCipher(shift int) *CaesarCipher {
	return &CaesarCipher{shift: shift}
}

func (c *CaesarCipher) shiftRune(r rune) rune {
	switch {
	case unicode.IsLetter(r):
		base := rune('a')
		if unicode.IsUpper(r) {
			base = rune('A')
		}
		shifted := (r - base + rune(c.shift)) % alphabetLength
		if shifted < 0 {
			shifted += alphabetLength
		}
		return base + shifted
	case unicode.IsDigit(r):
		shifted := (r - '0' + rune(c.shift)) % digitLength
		if shifted < 0 {
			shifted += digitLength
		}
		return '0' + shifted
	case r == '-':
		return r
	default:
		return r
	}
}

func (c *CaesarCipher) Encrypt(text string) string {
	var result strings.Builder
	for _, r := range text {
		result.WriteRune(c.shiftRune(r))
	}
	return result.String()
}

func (c *CaesarCipher) Decrypt(text string) string {
	decryptor := NewCaesarCipher(-c.shift)
	return decryptor.Encrypt(text)
}
