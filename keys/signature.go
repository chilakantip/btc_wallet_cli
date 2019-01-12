package keys

import (
	"github.com/chilakantip/btc_wallet_cli/secp256k1"
	"github.com/pkg/errors"
)

func (p *PrivateAddr) SignTransaction(trxMsg []byte) (signedTrx []byte, err error) {
	doubleHash := Sha2Sum(trxMsg)
	r, s, err := EcdsaSign(p.Key, doubleHash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign the trx.")
	}

	var sig secp256k1.Signature
	sig.R.Set(r)
	sig.S.Set(s)

	return sig.Bytes(), nil
}
