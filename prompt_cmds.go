package main

import (
	"github.com/c-bata/go-prompt"
)

const (
	cmdCreate          = "create"
	cmdRestoreFromSeed = "restore_from_seed"
	cmdImportkey       = "importkey"
	cmdAddress         = "address"
	cmdBalance         = "balance"
	cmdExportkey       = "exportkey"
	cmdQ               = "q"
)

func startCmds(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: cmdCreate, Description: "create a new wallet"},
		{Text: cmdRestoreFromSeed, Description: "restore wallet from seed"},
		{Text: cmdImportkey, Description: "import your wallet from WIP file"},
		{Text: cmdQ, Description: "exit"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func allCmds(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: cmdAddress, Description: "address to receive bitcoin"},
		{Text: cmdBalance, Description: "view your bitcoin wallet balance"},
		{Text: cmdExportkey, Description: "backup your private keys into WIP formate file"},
		{Text: cmdQ, Description: "exit"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
