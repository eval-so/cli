package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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
	jsonQuery, err := json.Marshal(e)
	if err != nil {
		printFatal("Error converting request into valid JSON.")
	}
	printDebug(fmt.Sprintf("JSON Query: %s", jsonQuery))

	query := bytes.NewBuffer(jsonQuery)

	parsedServer, err := url.Parse(*server)
	if err != nil {
		printFatal("Could not parse given server URL.")
	}

	resp, err := http.Post(parsedServer.String(), "application/json", query)
	if err != nil {
		printFatal("Could not successfully POST the query to the server.")
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		printFatal("Could not read the response from the server.")
	}

	return string(body), resp.StatusCode
}
