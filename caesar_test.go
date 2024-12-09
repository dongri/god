package god

import (
	"testing"
)

func TestCaesarCipher_Encrypt(t *testing.T) {
	tests := []struct {
		name     string
		shift    int
		input    string
		expected string
	}{
		{
			name:     "Basic alphabet encryption",
			shift:    3,
			input:    "HELLO",
			expected: "KHOOR",
		},
		{
			name:     "String with numbers encryption",
			shift:    3,
			input:    "HELLO-123",
			expected: "KHOOR-456",
		},
		{
			name:     "Mixed case encryption",
			shift:    3,
			input:    "Hello-World",
			expected: "Khoor-Zruog",
		},
		{
			name:     "Numbers only encryption",
			shift:    3,
			input:    "123456789",
			expected: "456789012",
		},
		{
			name:     "Encryption with negative shift",
			shift:    -3,
			input:    "HELLO-123",
			expected: "EBIIL-890",
		},
		{
			name:     "Encryption with large shift value",
			shift:    29, // 26 + 3
			input:    "HELLO",
			expected: "KHOOR",
		},
		{
			name:     "Encryption with special characters",
			shift:    3,
			input:    "HELLO! @#$%",
			expected: "KHOOR! @#$%",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cipher := NewCaesarCipher(tt.shift)
			if got := cipher.Encrypt(tt.input); got != tt.expected {
				t.Errorf("Encrypt() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCaesarCipher_Decrypt(t *testing.T) {
	tests := []struct {
		name     string
		shift    int
		input    string
		expected string
	}{
		{
			name:     "Basic alphabet decryption",
			shift:    3,
			input:    "KHOOR",
			expected: "HELLO",
		},
		{
			name:     "String with numbers decryption",
			shift:    3,
			input:    "KHOOR-456",
			expected: "HELLO-123",
		},
		{
			name:     "Mixed case decryption",
			shift:    3,
			input:    "Khoor-Zruog",
			expected: "Hello-World",
		},
		{
			name:     "Numbers only decryption",
			shift:    3,
			input:    "456789012",
			expected: "123456789",
		},
		{
			name:     "Decryption with negative shift",
			shift:    -3,
			input:    "EBIIL-890",
			expected: "HELLO-123",
		},
		{
			name:     "Decryption with large shift value",
			shift:    29, // 26 + 3
			input:    "KHOOR",
			expected: "HELLO",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cipher := NewCaesarCipher(tt.shift)
			if got := cipher.Decrypt(tt.input); got != tt.expected {
				t.Errorf("Decrypt() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCaesarCipher_EncryptDecryptRoundTrip(t *testing.T) {
	tests := []struct {
		name  string
		shift int
		input string
	}{
		{
			name:  "Basic round trip test",
			shift: 3,
			input: "HELLO-123",
		},
		{
			name:  "Mixed case round trip test",
			shift: 3,
			input: "Hello-World-123",
		},
		{
			name:  "Round trip test with special characters",
			shift: 3,
			input: "Hello! @#$% 123",
		},
		{
			name:  "Round trip test with negative shift",
			shift: -3,
			input: "HELLO-WORLD-123",
		},
		{
			name:  "Round trip test with large shift value",
			shift: 29,
			input: "HELLO-WORLD-123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cipher := NewCaesarCipher(tt.shift)
			encrypted := cipher.Encrypt(tt.input)
			decrypted := cipher.Decrypt(encrypted)
			if decrypted != tt.input {
				t.Errorf("Round trip failed: got %v, want %v", decrypted, tt.input)
			}
		})
	}
}
