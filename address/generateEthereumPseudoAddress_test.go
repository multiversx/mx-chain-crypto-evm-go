package address

import (
	"testing"

	"github.com/multiversx/mx-chain-core-evm-go/core"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePseudoEthereumAddress(t *testing.T) {
	_, err := GeneratePseudoAddress([]byte("multiversXAddress"), core.MVXAddressIdentifier, core.ETHAddressIdentifier)
	assert.NoError(t, err)
}
