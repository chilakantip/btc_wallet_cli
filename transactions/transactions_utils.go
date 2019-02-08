package transactions

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/chilakantip/btc_wallet_cli/utils"
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
	TxOutputN       uint32       `json:"tx_output_n"`
	Script          string       `json:"script"`
	Value           uint64       `json:"value"`
	ValueHex        string       `json:"value_hex"`
	Confirmations   uint64       `json:"confirmations"`
	rawTrx          bytes.Buffer `json:"rawtrx"`
	signature       []byte       `json:"signature"`
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
	//	if err = utils.ValidateBTCAddress(add); err != nil {
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
	hash160, err := utils.BTCAddHash160(s.btcAddress)
	if err != nil {
		return
	}
	// P2PKH transcation
	buf := new(bytes.Buffer)
	buf.WriteByte(OP_DUP)
	buf.WriteByte(OP_HASH160)
	buf.Write(utils.LittleEndianHex(int8(len(hash160))))
	buf.Write(hash160)
	buf.WriteByte(OP_EQUALVERIFY)
	buf.WriteByte(OP_CHECKSIG)

	s.scriptPubKey = append(s.scriptPubKey, utils.LittleEndianHex(uint8(buf.Len()))...)
	s.scriptPubKey = append(s.scriptPubKey, buf.Bytes()...)

	return nil
}

var resp = `{
    
    "unspent_outputs":[
    
        {
            "tx_hash":"f9c39e617156c02e923e4185a6b8324336d2aee0c9f28b1bd28508cde7703d3e",
            "tx_hash_big_endian":"3e3d70e7cd0885d21b8bf2c9e0aed2364332b8a685413e922ec05671619ec3f9",
            "tx_index":312340327,
            "tx_output_n": 2,
            "script":"76a914146a58f188a19c075d5f88ff4ae530a59bc1853888ac",
            "value": 546,
            "value_hex": "0222",
            "confirmations":59065
        },
      
        {
            "tx_hash":"40f1f69c5a63be63533b369065903ad8b0a5d341f804ce54d3026ef45dba93ba",
            "tx_hash_big_endian":"ba93ba5df46e02d354ce04f841d3a5b0d83a906590363b5363be635a9cf6f140",
            "tx_index":393703894,
            "tx_output_n": 1,
            "script":"76a914146a58f188a19c075d5f88ff4ae530a59bc1853888ac",
            "value": 118307,
            "value_hex": "28c3",
            "confirmations":6290
        }
      
    ]
}
`
