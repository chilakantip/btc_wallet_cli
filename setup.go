package main

import (
	"fmt"
	"os"
)

const (
	Success = iota
	SetupFailed
)

const (
	exportFileName = "btc_wallet"
)

func Quit() {
	fmt.Println("Quit command received; quitting application.")
	os.Exit(Success)
}

func abort() {
	os.Exit(SetupFailed)
}

var helpMsg = `:: CLI commands :: 
create: create a new wallet 
restore_from_seed: restore wallet from seed
importkey: import your wallet from WI file 
address: Address to receive bitcoin
balance: view bitcoin wallet balance
exportkey: backup your private keys into WIP formate file 
q: exit

please enter the command(hit tab for )
`
