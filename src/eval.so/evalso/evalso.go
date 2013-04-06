package main

import (
	"eval.so/api"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	action   = flag.String("action", "evaluate", "What should I do?")
	language = flag.String("language", "", "Language to evaluate.")
	debug    = flag.Bool("debug", false, "Print debugging output?")
	server   = flag.String("server", "http://eval.so/jsontest", "API server to evaluate against.")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: evalso [flags]\n")
	fmt.Fprintf(os.Stderr, "Defaults are:\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func printDebug(message string) {
	if *debug {
		log.Printf("[\x1b[33;1mDEBUG\x1b[0m] %s\n", message)
	}
}

func printFatal(message string) {
	log.Fatal(fmt.Sprintf("[\x1b[31;1mERROR\x1b[0m] %s\n", message))
}

func printOkay(message string) {
	log.Printf("[\x1b[32;1mOKAY\x1b[0m] %s\n", message)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
	}

	printDebug("Requesting evaluation...")
	code, _ := ioutil.ReadAll(os.Stdin)
	e := evaluate.Evaluation{
		Language: *language,
		Code:     string(code),
	}

	result, statusCode := evaluate.Evaluate(*server, e)

	if statusCode != 200 {
		printFatal(fmt.Sprintf("HTTP Code %d", statusCode))
	} else {
		fmt.Println(string(result))
	}
}
