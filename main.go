package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tendermint/tendermint/libs/bech32"
)

var flagHelp bool

func main() {
	flag.BoolVar(&flagHelp, "help", false, "Show help")
	flag.Parse()

	if flagHelp {
		fmt.Println(`
		Convert HRP of a bech32 encoded string and regenerate the checksum
		1st argument desired prefix. 
		2nd argument is the bech32 to convert
		`)
		return
	}
	targetHRP := os.Args[1]
	bech32str := os.Args[2]

	if targetHRP == "" {
		fmt.Fprintf(os.Stderr, "Missing target HRP\n")
		return
	}
	if bech32str == "" {
		fmt.Fprintf(os.Stderr, "Missing target bech32 phrase\n")
		return
	}

	_, bz, err := bech32.DecodeAndConvert(bech32str)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding bech32: %s\n", err)
		return
	}
	updatedHRP, err := bech32.ConvertAndEncode(targetHRP, bz)
	fmt.Println(updatedHRP)
}
