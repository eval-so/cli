package evaluate

import (
	"bytes"
	"encoding/json"
	"eval.so/util"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Evaluation struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

type EvaluationResult struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	Walltime int `json:"wallTime"`
	Exitcode int `json:"exitCode"`
}

func Evaluate(server string, e Evaluation) (EvaluationResult, int) {
	jsonQuery, err := json.Marshal(e)
	if err != nil {
		util.PrintFatal("Error converting request into valid JSON.")
	}
	//PrintDebug(fmt.Sprintf("JSON Query: %s", jsonQuery))

	query := bytes.NewBuffer(jsonQuery)

	parsedServer, err := url.Parse(server)
	if err != nil {
		util.PrintFatal("Could not parse given server URL.")
	}

	resp, err := http.Post(parsedServer.String(), "application/json", query)
	if err != nil {
		util.PrintFatal("Could not successfully POST the query to the server.")
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		util.PrintFatal("Could not read the response from the server.")
	}

	var result EvaluationResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		util.PrintFatal("Error decoding JSON response from the server.")
	}

	return result, resp.StatusCode
}
