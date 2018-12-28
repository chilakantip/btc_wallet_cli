package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/btcsuite/btcutil"
	"github.com/c-bata/go-prompt"
	"github.com/chilakantip/btc_wallet_cli/api"
	"github.com/chilakantip/btc_wallet_cli/keys"
)

func main() {
	fmt.Println(helpMsg)
	var pk *keys.PrivateAddr

	done := false
	for !done {
		opt := prompt.Input("> ", startCmds)

		switch strings.TrimSpace(opt) {
		case cmdCreate, cmdRestoreFromSeed:
			seed, err := getSeedFromCli()
			if err != nil {
				fmt.Println("failed to create keys, try again")
				break
			}

			pk, err = keys.SetupKeysFromSeed(seed)
			if err != nil {
				fmt.Println("\nfailed to create keys, try again")
				abort()
			}
			fmt.Println("\nsetup keys success")
			fmt.Println("your BTC wallet address: ", pk.Address)
			done = true
			break
		case cmdImportkey:
			fmt.Print("Enter WIF file name: ")
			file, err := getLineFromCmdLine()
			if err != nil {
				fmt.Println("\nfailed to read file name, try again")
				break
			}

			pk = keys.GetKeyTemplate()
			if err := pk.ImportWIF(strings.TrimSpace(file)); err != nil {
				fmt.Println("failed to import the WIF file")
				abort()
			}
			fmt.Println("wallet import success")
			fmt.Println("your BTC wallet address: ", pk.Address)
			done = true
			break
		case cmdQ:
			abortQuit()
		default:
			fmt.Println("invalid option, please select correct option")
		}
	}

	for {
		opt := prompt.Input("> ", allCmds)
		switch strings.TrimSpace(opt) {
		case cmdAddress:
			if pk.Address == "" {
				fmt.Println("failed to get BTC address")
				break
			}
			fmt.Println("your BTC wallet address: ", pk.Address)
			break

		case cmdExportkey:
			if err := pk.ExportWIF(); err != nil {
				fmt.Println("failed to export private key")
				break
			}

			fileName := fmt.Sprintf("%s_%s.wif", exportFileName, time.Now().Format("2006_01_02_150405"))
			if err := ioutil.WriteFile(fileName, []byte(pk.WIF), 0400); err != nil {
				fmt.Println("failed to export private key")
				break
			}
			fmt.Println("exported private key to", exportFileName, "file")
			break
		case cmdBalance:
			bal, err := api.GetBTCBalance(pk.Address)
			if err != nil {
				fmt.Println("failed to get balance")
			}

			balInBTC, _ := btcutil.NewAmount(bal.Balance / btcutil.SatoshiPerBitcoin)
			fmt.Println("your BTC balance:", balInBTC.Format(btcutil.AmountBTC))
			break
		case cmdQ:
			abortQuit()
		default:
			fmt.Println("invalid option, please select correct option")
		}
	}
}
