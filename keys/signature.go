package keys

import (
	"bytes"

	"github.com/chilakantip/btc_wallet_cli/utils"

	"github.com/chilakantip/btc_wallet_cli/secp256k1"
	"github.com/pkg/errors"
)

func (p *PrivateAddr) SignTransaction(trxMsg []byte) (signedTrx []byte, err error) {
	doubleHash := utils.DoubleHashH(trxMsg)
	r, s, err := EcdsaSign(p.Key, doubleHash.CloneBytes())
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign the trx.")
	}

	var sig secp256k1.Signature
	sig.R.Set(r)
	sig.S.Set(s)

	return sig.Bytes(), nil
}

func (p *PrivateAddr) CalcScriptSig(sig []byte) (ssig []byte) {
	ssBuf := new(bytes.Buffer)
	var scriptSigVer = byte(0x01)

	sig = append(sig, scriptSigVer)
	sigLen := utils.LittleEndianHex(uint8(len(sig)))
	sig = append(sigLen, sig...)

	pubklen := utils.LittleEndianHex(uint8(len(p.Pubkey)))
	pubkAndlen := append(pubklen, p.Pubkey...)

	ssBuf.WriteByte(sigLen[0] + pubklen[0])
	ssBuf.Write(sig)
	ssBuf.Write(pubkAndlen)

	return ssBuf.Bytes()
}
