package main

import (
	"crypto/aes"
	"flag"
	"fmt"
	"giwifi-mix/util"
	"os"
)

const (
	defaultKey = "1234567887654321"
)

var (
	key, iv, text string
	ishelp        bool
)

func initFlag() {
	flag.StringVar(&text, "t", "", "the text to be encrypted")
	flag.StringVar(&iv, "i", "", "iv (initialization vector)")
	flag.StringVar(&key, "k", defaultKey, "the key")
	flag.BoolVar(&ishelp, "h", false, "print this help menu")

	flag.Parse()

	// check the avg
	if text == "" || iv == "" {
		flag.CommandLine.Usage()
		os.Exit(1)
	}

	if len(key) != 16 || len(iv) != 16 {
		fmt.Printf("error")
		os.Exit(1)
	}

}

func main() {
	initFlag()
	result, err := util.Crypto(text, key, iv, aes.BlockSize)
	if err != nil {
		fmt.Printf("error")
		os.Exit(1)
	}
	fmt.Printf(result)
}