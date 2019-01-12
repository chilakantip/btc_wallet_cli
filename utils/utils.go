package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/btcsuite/btcutil/base58"
)

func LittleEndianHex(s interface{}) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, s)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	return buf.Bytes()
}

func HexAndReverseStr(s string) (hexArr []byte, err error) {
	hexArr, err = hex.DecodeString(s)
	if err != nil {
		return
	}

	return reverse(hexArr), nil
}

func reverse(nums []byte) []byte {
	newnums := make([]byte, len(nums))
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		newnums[i], newnums[j] = nums[j], nums[i]
	}
	return newnums
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

func ValidatePrivateKeyWIF(wif []byte) error {
	if wif[0] != 0x80 {
		return base58.ErrInvalidFormat
	}
	if !bytes.Equal(wif[33:37], Checksum(wif[0:33])) {
		return base58.ErrChecksum
	}

	return nil
}

func Checksum(in []byte) (sum []byte) {
	sh := Sha2Sum(in)
	sum = make([]byte, 4)
	copy(sum, sh[:4])
	return
}
