package god

import (
	"encoding/hex"
	"math/big"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeSignature(t *testing.T) {
	t.Run("Decode Signature(Swap)", func(t *testing.T) {
		signature := "Swap(int256,int256,uint160,uint128,int24)"
		dataHex := "0x000000000000000000000000000000000000000000000003cb71f51fc5580000ffffffffffffffffffffffffffffffffffffffffffffffffffd31ba748ba48cf000000000000000000000000000000000000000003712b8d4030f698cef2e00b000000000000000000000000000000000000000000000021b81b8b36054a6c8bfffffffffffffffffffffffffffffffffffffffffffffffffffffffffffeaf55"

		expected := []any{
			strToBigInt("70000000000000000000"),
			strToBigInt("-12635968658323249"),
			strToBigInt("1065269314455634044184092683"),
			strToBigInt("622008904623898389643"),
			strToBigInt("-86187"),
		}

		dataBytes, err := hex.DecodeString(strings.TrimPrefix(dataHex, "0x"))
		if err != nil {
			t.Fatalf("Failed to decode hex string: %v", err)
		}

		decodedArgs, err := DecodeSignature(signature, dataBytes)
		require.NoError(t, err)
		assert.Len(t, decodedArgs, len(expected))

		assert.Equal(t, decodedArgs[0], expected[0])
		assert.Equal(t, decodedArgs[1], expected[1])
		assert.Equal(t, decodedArgs[2], expected[2])
		assert.Equal(t, decodedArgs[3], expected[3])
		assert.Equal(t, decodedArgs[4], expected[4])

	})
}

func strToBigInt(s string) *big.Int {
	bi := new(big.Int)
	bi.SetString(s, 10)
	return bi
}
