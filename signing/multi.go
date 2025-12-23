package signing

import (
	"fmt"

	mvxCrypto "github.com/multiversx/mx-chain-crypto-go"

	"github.com/multiversx/mx-chain-crypto-evm-go"
)

var _ mvxCrypto.KeyGenerator = (*MultiKeyGenerator)(nil)

var _ mvxCrypto.SingleSigner = (*MultiSingleSigner)(nil)

// MultiKeyGenerator handles a set of key generators - selector based
type MultiKeyGenerator struct {
	mvxCrypto.KeyGenerator
	MainSelector    string
	OtherGenerators map[string]mvxCrypto.KeyGenerator
}

// MultiSingleSigner handles a set of single singers - selector based
type MultiSingleSigner struct {
	mvxCrypto.SingleSigner
	MainSelector string
	OtherSigners map[string]mvxCrypto.SingleSigner
}

// ChooseKeyGenerator returns the key generator to be used based on the given selector
func (kg *MultiKeyGenerator) ChooseKeyGenerator(selector string) (mvxCrypto.KeyGenerator, error) {
	return chooseForSelector(selector, kg.MainSelector, kg.KeyGenerator, kg.OtherGenerators)
}

// ChooseSingleSigner returns the single signer to be used based on the given selector
func (ss *MultiSingleSigner) ChooseSingleSigner(selector string) (mvxCrypto.SingleSigner, error) {
	return chooseForSelector(selector, ss.MainSelector, ss.SingleSigner, ss.OtherSigners)
}

func chooseForSelector[T interface{}](selector string, mainSelector string, mainItem T, otherItems map[string]T) (T, error) {
	if selector != mainSelector {
		otherItem, found := otherItems[selector]
		if !found {
			return otherItem, fmt.Errorf("%w: %v", crypto.ErrImplementationNotDefinedForSelector, selector)
		}
		return otherItem, nil
	}
	return mainItem, nil
}
