package transactions

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/chilakantip/btc_wallet_cli/keys"
)

var (
	OP_DUP         = byte(0x76)
	OP_HASH160     = byte(0xa9)
	OP_EQUALVERIFY = byte(0x88)
	OP_CHECKSIG    = byte(0xac)
)

// API ref https://www.blockchain.com/api/blockchain_api
const (
	bcInfoGetUTXO  = "https://blockchain.info/unspent?active=%s"
	bcInfoGetBlock = "https://blockchain.info/rawblock/%s"
)

type unspentOutputs struct {
	TxHash          string       `json:"tx_hash"`
	TxHashBigEndian string       `json:"tx_hash_big_endian"`
	TxIndex         uint64       `json:"tx_index"`
	TxOutputN       uint64       `json:"tx_output_n"`
	Script          string       `json:"script"`
	Value           uint64       `json:"value"`
	ValueHex        string       `json:"value_hex"`
	Confirmations   uint64       `json:"confirmations"`
	rawTrx          bytes.Buffer `json:"rawtrx"`
	scriptSig       []byte       `json:"script_sig"`
}

type P2ScriptPubKey struct {
	amount       uint64
	btcAddress   string
	scriptPubKey []byte
}

type P2PKHTrx struct {
	version      []byte
	vin          []unspentOutputs
	sequence     []byte
	vout         []P2ScriptPubKey
	lockTime     []byte
	hashCodeType []byte
	signedTrx    bytes.Buffer
}

type UTXOResp struct {
	UTXO []unspentOutputs `json:"unspent_outputs"`
}

func (s *UTXOResp) string() {
	for _, v := range s.UTXO {
		fmt.Println(fmt.Sprintf("%d Satoshi with tx hash %s", v.Value, v.TxHash))
	}
}

func getUTXO(add string) (utxo *UTXOResp, err error) {
	utxo = new(UTXOResp)
	//	if err = keys.ValidateBTCAddress(add); err != nil {
	//		return
	//	}

	//	resp, err := resty.R().Get(fmt.Sprintf(bcInfoGetUTXO, add))
	//	if err != nil {
	//		return utxo, errors.Wrap(err, "failed to get UTXO from blockchain.info")
	//	}

	//	err = json.Unmarshal(resp.Body(), &utxo)

	err = json.Unmarshal([]byte(resp), &utxo)
	return

}

//four-byte version field
//four-byte field denoting the sequence. This is currently always set to 0xffffffff:
//four-byte "lock time" field: 00000000
//four-byte "hash code type"

func getP2PKHTrxVer1Template() *P2PKHTrx {
	p2pkh := new(P2PKHTrx)

	p2pkh.version, p2pkh.sequence, p2pkh.lockTime, p2pkh.hashCodeType =
		make([]byte, 4), make([]byte, 4), make([]byte, 4), make([]byte, 4)
	copy(p2pkh.version, []byte{0x01, 0x00, 0x00, 0x00})
	copy(p2pkh.sequence, []byte{0xff, 0xff, 0xff, 0xff})
	copy(p2pkh.lockTime, []byte{0x00, 0x00, 0x00, 0x00})
	copy(p2pkh.hashCodeType, []byte{0x01, 0x00, 0x00, 0x00})
	return p2pkh
}

func (s *P2ScriptPubKey) calcScriptPubKey() (err error) {
	hash160, err := keys.BTCAddHash160(s.btcAddress)
	if err != nil {
		return
	}
	// P2PKH transcation
	buf := new(bytes.Buffer)
	buf.WriteByte(OP_DUP)
	buf.WriteByte(OP_HASH160)
	buf.Write(toLittleEndianHex(int8(len(hash160))))
	buf.Write(hash160)
	buf.WriteByte(OP_EQUALVERIFY)
	buf.WriteByte(OP_CHECKSIG)

	s.scriptPubKey = append(s.scriptPubKey, toLittleEndianHex(uint8(buf.Len()))...)
	s.scriptPubKey = append(s.scriptPubKey, buf.Bytes()...)

	return nil
}
