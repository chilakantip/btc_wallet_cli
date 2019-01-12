package transactions

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

func toLittleEndianHex(s interface{}) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, s)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	return buf.Bytes()
}

func hexAndReverseStr(s string) (hexArr []byte, err error) {
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

var resp = `{
    
    "unspent_outputs":[
    
        {
            "tx_hash":"5adbb9fb73d8240a8385ba03293d27811239279824250e2aeb9cde6048a0679b",
            "tx_hash_big_endian":"9b67a04860de9ceb2a0e25249827391281273d2903ba85830a24d873fbb9db5a",
            "tx_index":383803601,
            "tx_output_n": 783,
            "script":"76a914b291bccae6b43a3ab9693b530605cda8c741a1dc88ac",
            "value": 200,
            "value_hex": "0378",
            "confirmations":10217
        },
      
        {
            "tx_hash":"12e40cb1cebfb5667fc39cda71f3e0c63dfc5d299f354a6b26365dbfa3e025b7",
            "tx_hash_big_endian":"b725e0a3bf5d36266b4a359f295dfc3dc6e0f371da9cc37f66b5bfceb10ce412",
            "tx_index":383864898,
            "tx_output_n": 628,
            "script":"76a914b291bccae6b43a3ab9693b530605cda8c741a1dc88ac",
            "value": 888,
            "value_hex": "0378",
            "confirmations":10174
        },
      
        {
            "tx_hash":"ba36150e1fb7a9f6a37040873401c5be8819acd06f7d4e8b2c86a8205c481859",
            "tx_hash_big_endian":"5918485c20a8862c8b4e7d6fd0ac1988bec50134874070a3f6a9b71f0e1536ba",
            "tx_index":403839123,
            "tx_output_n": 1,
            "script":"76a914b291bccae6b43a3ab9693b530605cda8c741a1dc88ac",
            "value": 500,
            "value_hex": "0ff47900",
            "confirmations":27
        },
      
        {
            "tx_hash":"9b8cb418bb138f699c78d8662368998bef7fd1593967e94248197fa58e7cf0b3",
            "tx_hash_big_endian":"b3f07c8ea57f194842e9673959d17fef8b99682366d8789c698f13bb18b48c9b",
            "tx_index":403875035,
            "tx_output_n": 1,
            "script":"76a914b291bccae6b43a3ab9693b530605cda8c741a1dc88ac",
            "value": 400,
            "value_hex": "19944f30",
            "confirmations":12
        },
      
        {
            "tx_hash":"e4df6d27c71e847cd5a36a3ff84776150ba5b9c0f12a18ebc643099f0607d982",
            "tx_hash_big_endian":"82d907069f0943c6eb182af1c0b9a50b157647f83f6aa3d57c841ec7276ddfe4",
            "tx_index":403888023,
            "tx_output_n": 1,
            "script":"76a914b291bccae6b43a3ab9693b530605cda8c741a1dc88ac",
            "value": 300,
            "value_hex": "12ad5770",
            "confirmations":0
        },
      
        {
            "tx_hash":"93c54e53749f258b509c38da51a36183ed4eb693a1fa4b8609b39adf51a20788",
            "tx_hash_big_endian":"8807a251df9ab309864bfaa193b64eed8361a351da389c508b259f74534ec593",
            "tx_index":403902910,
            "tx_output_n": 0,
            "script":"76a914b291bccae6b43a3ab9693b530605cda8c741a1dc88ac",
            "value": 800,
            "value_hex": "016caf60",
            "confirmations":0
        },
      
        {
            "tx_hash":"cebd6e7c62a5cf8780f91db15d9c712c4a93a532a0138c6bd78527b5f2bfbcea",
            "tx_hash_big_endian":"eabcbff2b52785d76b8c13a032a5934a2c719c5db11df98087cfa5627c6ebdce",
            "tx_index":403902430,
            "tx_output_n": 0,
            "script":"76a914b291bccae6b43a3ab9693b530605cda8c741a1dc88ac",
            "value": 900,
            "value_hex": "05e4ca00",
            "confirmations":0
        }
      
    ]
}
`
