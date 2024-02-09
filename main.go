package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	args := parseArgs()

	var result []byte
	var err error
	switch args.mode {
	case "encrypt":
		result, err = Encrypt(args.input, args.key)
		if err != nil {
			fmt.Println("error encrypting:", err)
			os.Exit(1)
		}
	case "decrypt":
		result, err = Decrypt(args.input, args.key)
		if err != nil {
			fmt.Println("error decrypting:", err)
			os.Exit(1)
		}
	}

	err = os.WriteFile(args.outputFile, result, fs.ModePerm)
	if err != nil {
		fmt.Println("error writing file:", err)
		os.Exit(1)
	}
}
