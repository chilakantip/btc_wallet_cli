package transactions

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"sort"

	"github.com/chilakantip/btc_wallet_cli/keys"
	"github.com/chilakantip/btc_wallet_cli/utils"
	"github.com/pkg/errors"
)

func MakeTxMsg(pk *keys.PrivateAddr, toAdd string, satoshi uint64) (err error) {
	utxoResp, err := getUTXO(pk.Address)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to get UTXO for address: %s", pk.Address))
	}

	// please consider the fee; not added yet
	seldUTXO, change, err := selectTrxToSpend(utxoResp.UTXO, satoshi)
	if err != nil {
		return errors.Wrap(err, "failed to select UTXO to spend amount")
	}

	trxTemp := getP2PKHTrxVer1Template()

	//create 2 outputs, one to send address and take change back your address
	sendVout1 := P2ScriptPubKey{amount: satoshi, btcAddress: toAdd}
	trxTemp.vout = append(trxTemp.vout, sendVout1)
	if change > uint64(0) {
		changeBackVout2 := P2ScriptPubKey{amount: change, btcAddress: pk.Address}
		trxTemp.vout = append(trxTemp.vout, changeBackVout2)
	}

	// TODO: rise error when no change(fee) else trx will never be confirmed

	//create the vouts
	voutBuf := new(bytes.Buffer)
	voutBuf.Write(utils.LittleEndianHex(uint8(len(trxTemp.vout))))
	for _, singleVout := range trxTemp.vout {
		voutBuf.Write(utils.LittleEndianHex(uint64(singleVout.amount)))
		if err = singleVout.calcScriptPubKey(); err != nil {
			return errors.Wrap(err, "failed to calculate the ScriptPub")
		}
		voutBuf.Write(singleVout.scriptPubKey)
	}

	// generat the Vins
	vinBuf := new(bytes.Buffer)
	vinBuf.Write(utils.LittleEndianHex(uint8(len(seldUTXO))))
	for _, singleUtxo := range seldUTXO {
		trxHash, err_ := utils.HexAndReverseStr(singleUtxo.TxHash)
		if err_ != nil {
			return err_
		}
		vinBuf.Write(trxHash)
		vinBuf.Write(utils.LittleEndianHex(uint64(singleUtxo.TxOutputN)))

	}

	// generat scriptSig
	scriptSigBuf := new(bytes.Buffer)
	for _, singleUtxo := range seldUTXO {
		trxScript, err_ := hex.DecodeString(singleUtxo.Script)
		if err_ != nil {
			return
		}

		signBuf := new(bytes.Buffer)

		signBuf.Write(trxTemp.version)
		signBuf.Write(vinBuf.Bytes())

		signBuf.Write(utils.LittleEndianHex(uint8(len(trxScript))))
		signBuf.Write(trxScript)
		signBuf.Write(trxTemp.sequence)

		signBuf.Write(voutBuf.Bytes())

		signBuf.Write(trxTemp.lockTime)
		signBuf.Write(trxTemp.hashCodeType)

		//now msg is ready for signature

		sig, err := pk.SignTransaction(signBuf.Bytes())
		if err != nil {
			return errors.Wrap(err, "failed to sign the transaction")
		}
		copy(singleUtxo.signature, sig)
		copy(singleUtxo.scriptSig, pk.CalcScriptSig(sig))

		scriptSigBuf.Write(singleUtxo.scriptSig)
		scriptSigBuf.Write(trxTemp.sequence)
	}

	//construct the final transaction
	finalSigTrx := new(bytes.Buffer)

	finalSigTrx.Write(trxTemp.version)
	finalSigTrx.Write(vinBuf.Bytes())
	finalSigTrx.Write(scriptSigBuf.Bytes())
	finalSigTrx.Write(voutBuf.Bytes())
	finalSigTrx.Write(trxTemp.lockTime)
	finalSigTrx.Write(trxTemp.hashCodeType)

	fmt.Println(fmt.Sprintf("%x", finalSigTrx.Bytes()))
	return nil
}

func selectTrxToSpend(utxo []unspentOutputs, minAmount uint64) (seldUTXO []unspentOutputs, change uint64, err error) {
	if len(utxo) == 0 || minAmount == 0 {
		return nil, 0, fmt.Errorf("no utxo input")
	}

	var lesser []unspentOutputs
	var greater []unspentOutputs

	for _, v := range utxo {
		if v.Value < minAmount {
			lesser = append(lesser, v)
		} else {
			greater = append(greater, v)
		}
	}

	if len(greater) > 0 {
		sort.Slice(greater, func(i, j int) bool {
			return greater[i].Value < greater[j].Value
		})
		seldUTXO = append(seldUTXO, greater[0])

		return seldUTXO, uint64(greater[0].Value - minAmount), nil

	}

	sort.Slice(lesser, func(i, j int) bool {
		return lesser[i].Value > lesser[j].Value
	})

	acc := uint64(0)
	for _, v := range lesser {
		seldUTXO = append(seldUTXO, v)
		acc = acc + v.Value
		if acc >= minAmount {
			return seldUTXO, uint64(acc - minAmount), nil
		}
	}

	return nil, 0, fmt.Errorf("no matched utxo found")
}
