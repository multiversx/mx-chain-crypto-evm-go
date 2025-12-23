package address

import (
	"testing"

	"github.com/multiversx/mx-chain-core-evm-go/core"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePseudoMultiversXAddress(t *testing.T) {
	_, err := GeneratePseudoAddress([]byte("ethereumAddress"), core.ETHAddressIdentifier, core.MVXAddressIdentifier)
	assert.NoError(t, err)
}
