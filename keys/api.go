package keys

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/chilakantip/btc_wallet_cli/utils"
	"github.com/pkg/errors"
	"gopkg.in/resty.v1"
)

type BTCAddInfo struct {
	Address         string
	Balance         float64
	Unit            string
	No_tx           int64
	Totail_received int64
}

// API ref https://www.blockchain.com/api/blockchain_api
const (
	bcInfoCheckBalance = "https://blockchain.info/balance?active=%s"

//	bcInfoGetUTXO      = "https://blockchain.info/unspent?active=%s"
)

func GetBTCBalance(add string) (btc *BTCAddInfo, err error) {
	btc = new(BTCAddInfo)
	if err = utils.ValidateBTCAddress(add); err != nil {
		return
	}

	resp, err := resty.R().Get(fmt.Sprintf(bcInfoCheckBalance, add))
	if err != nil {
		return btc, errors.Wrap(err, "failed to get balance from blockchain.info")
	}

	var btcInfo interface{}
	if err = json.Unmarshal(resp.Body(), &btcInfo); err != nil {
		return
	}

	btcAddInfo := btcInfo.(map[string]interface{})
	balMap := btcAddInfo[add]
	v := reflect.ValueOf(balMap)
	if v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			if key.String() == "final_balance" {
				btc.Balance = v.MapIndex(key).Interface().(float64)
				btc.Address = add
				btc.Unit = "Satoshi"
			}
		}
	}

	return
}
