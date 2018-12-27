package keys

type PrivateAddr struct {
	Version byte
	Key     []byte
	WIF     string
	*BtcAddr
}

type BtcAddr struct {
	Version  byte
	Hash160  []byte
	Checksum []byte
	Pubkey   []byte
	Address  string //base58 encoded
}

func GetKeyTemplate() *PrivateAddr {
	pk := new(PrivateAddr)
	pk.BtcAddr = new(BtcAddr)
	pk.Version = 0x80
	pk.Key = make([]byte, 32)

	pk.BtcAddr.Version = 0x00
	pk.Hash160 = make([]byte, 20)
	pk.Pubkey = make([]byte, 65) // uncompressed public key
	pk.Checksum = make([]byte, 4)

	return pk
}
