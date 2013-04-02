package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	action   = flag.String("action", "evaluate", "What should I do?")
	language = flag.String("language", "", "Language to evaluate.")
	debug    = flag.Bool("debug", false, "Print debugging output?")
	server   = flag.String("server", "http://eval.gd", "API server to evaluate against.")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: evalgd [flags]\n")
	fmt.Fprintf(os.Stderr, "Defaults are:\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func printDebug(message string)  {
	if *debug {
		fmt.Printf("[\x1b[33;1mDEBUG\x1b[0m] %s\n", message)
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
	}

	printDebug("Requesting evaluation...")
	code, _ := ioutil.ReadAll(os.Stdin)
	e := Evaluation{
		Language: *language,
		Code:     string(code),
	}

	result, statusCode := evaluate(e)

	if statusCode != 200 {
		fmt.Printf("[\x1b[31;1mError\x1b[0m] HTTP Code %d\n", statusCode)
	} else {
		fmt.Println(string(result))
	}
}

type Evaluation struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

func evaluate(e Evaluation) (string, int) {
	jsonQuery, _ := json.Marshal(e)
	query := bytes.NewBuffer(jsonQuery)
	resp, _ := http.Post(
		fmt.Sprintf("%s/jsontest", *server),
		"application/json",
		query)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), resp.StatusCode
}
