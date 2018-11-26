package main

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/gotomgo/bingSpellCheck"
)

const key = "<INSERT YOUR API KEY HERE>"

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Provide some text to spell/grammar check on the command line")
		os.Exit(1)
	}

	client := bingSpellCheck.NewClient(key)

	spellCheck, err := client.SpellCheck(os.Args[1])
	if err != nil {
		fmt.Println(err)
	} else {
		spew.Dump(spellCheck)
	}

	correctedText, err := bingSpellCheck.BuildAutoCorrectedText(os.Args[1], spellCheck)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(correctedText)
	}

}
