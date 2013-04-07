package main

import (
	"eval.so/api"
	"eval.so/util"
	"flag"
	"fmt"
	"io/ioutil"
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

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
	}

	util.PrintDebug("Requesting evaluation...")
	code, _ := ioutil.ReadAll(os.Stdin)
	e := evaluate.Evaluation{
		Language: *language,
		Code:     string(code),
	}

	result, statusCode := evaluate.Evaluate(*server, e)

	if statusCode != 200 {
		util.PrintFatal(fmt.Sprintf("HTTP Code %d", statusCode))
	} else {
		fmt.Println("[\x1b[32;1mSTDOUT\x1b[0m]")
		fmt.Println(result.Stdout)

		fmt.Println("[\x1b[32;1mSTDERR\x1b[0m]")
		fmt.Println(result.Stderr)

		fmt.Println("[\x1b[32;1mWall Time\x1b[0m]")
		fmt.Printf("%dms\n", result.Walltime)
	}
}
