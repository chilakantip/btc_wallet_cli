package transactions

type UTXO struct {
	UnspentOutputs []struct {
		TxHash          string `json:"tx_hash"`
		TxHashBigEndian string `json:"tx_hash_big_endian"`
		TxIndex         int    `json:"tx_index"`
		TxOutputN       int    `json:"tx_output_n"`
		Script          string `json:"script"`
		Value           int    `json:"value"`
		ValueHex        string `json:"value_hex"`
		Confirmations   int    `json:"confirmations"`
	} `json:"unspent_outputs"`
}
