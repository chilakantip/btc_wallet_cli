package keys

import (
	"fmt"
	"io/ioutil"

	"github.com/chilakantip/btc_wallet_cli/utils"

	"github.com/btcsuite/btcutil/base58"
	"github.com/chilakantip/btc_wallet_cli/secp256k1"
)

func SetupKeysFromSeed(seed string) (pk *PrivateAddr, err error) {
	pk = GetKeyTemplate()
	utils.ShaHash([]byte(seed), pk.Key)
	err = pk.generateBTCAddFromPrivKey()

	return
}

func (pk *PrivateAddr) ExportWIF() (err error) {
	var wif [37]byte    //1 byte version, 32 bytes private key and 4 bytes checksum
	wif[0] = pk.Version //prefix 5 for WIP
	copy(wif[1:33], pk.Key)
	copy(wif[33:37], utils.Checksum(wif[0:33]))
	pk.WIF = base58.Encode(wif[:])

	return nil
}

func (pk *PrivateAddr) ImportWIF(file string) (err error) {
	key, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	wif := base58.Decode(string(key))
	if err = utils.ValidatePrivateKeyWIF(wif); err != nil {
		return
	}

	copy(pk.Key, wif[1:33])
	return pk.generateBTCAddFromPrivKey()

}

func (pk *PrivateAddr) generateBTCAddFromPrivKey() error {
	if !secp256k1.BaseMultiply(pk.Key, pk.Pubkey) {
		return fmt.Errorf("fail to generate public key")
	}
	utils.RimpHash(pk.Pubkey, pk.Hash160)

	var ad [25]byte
	ad[0] = pk.BtcAddr.Version
	copy(ad[1:21], pk.Hash160)
	copy(ad[21:25], utils.Checksum(ad[0:21]))
	pk.Address = base58.Encode(ad[:])

	return nil
}
