package keys

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"github.com/chilakantip/btc_wallet_cli/secp256k1"
)

func SetupKeysFromSeed(seed string) (pk *PrivateAddr, err error) {
	pk = GetKeyTemplate()
	ShaHash([]byte(seed), pk.Key)
	err = pk.generateBTCAddFromPrivKey()

	return
}

func (pk *PrivateAddr) ExportWIF() (err error) {
	var wif [37]byte    //1 byte version, 32 bytes private key and 4 bytes checksum
	wif[0] = pk.Version //prefix 5 for WIP
	copy(wif[1:33], pk.Key)
	copy(wif[33:37], Checksum(wif[0:33]))
	pk.WIF = base58.Encode(wif[:])

	return nil
}

func (pk *PrivateAddr) ImportWIF(file string) (err error) {
	key, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	wif := base58.Decode(string(key))
	if err = validatePrivateKeyWIF(wif); err != nil {
		return
	}

	copy(pk.Key, wif[1:33])
	return pk.generateBTCAddFromPrivKey()

}

func validatePrivateKeyWIF(wif []byte) error {
	if wif[0] != 0x80 {
		return base58.ErrInvalidFormat
	}
	if !bytes.Equal(wif[33:37], Checksum(wif[0:33])) {
		return base58.ErrChecksum
	}

	return nil
}

func (pk *PrivateAddr) generateBTCAddFromPrivKey() error {
	if !secp256k1.BaseMultiply(pk.Key, pk.Pubkey) {
		return fmt.Errorf("fail to generate public key")
	}
	RimpHash(pk.Pubkey, pk.Hash160)

	var ad [25]byte
	ad[0] = pk.BtcAddr.Version
	copy(ad[1:21], pk.Hash160)
	copy(ad[21:25], Checksum(ad[0:21]))
	pk.Address = base58.Encode(ad[:])

	return nil
}

func ValidateBTCAddress(add string) error {
	if strings.TrimSpace(add) == "" {
		return fmt.Errorf("empty bitcoin address")
	}

	addBytes := base58.Decode(add)
	if addBytes[0] != 0x00 {
		return base58.ErrInvalidFormat
	}
	if !bytes.Equal(addBytes[21:25], Checksum(addBytes[0:21])) {
		return base58.ErrChecksum
	}

	return nil
}

func BTCAddHash160(add string) (hash160 []byte, err error) {
	if err = ValidateBTCAddress(add); err != nil {
		return
	}

	hash160 = make([]byte, 20)
	buf := base58.Decode(add)
	copy(hash160, buf[1:21])
	return
}
