package util

import (
	"fmt"
	"log"
)


func PrintDebug(message string) {
	log.Printf("[\x1b[33;1mDEBUG\x1b[0m] %s\n", message)
}

func PrintFatal(message string) {
	log.Fatal(fmt.Sprintf("[\x1b[31;1mERROR\x1b[0m] %s\n", message))
}

func PrintOkay(message string) {
	log.Printf("[\x1b[32;1mOKAY\x1b[0m] %s\n", message)
}
