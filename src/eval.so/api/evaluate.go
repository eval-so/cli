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

func Evaluate(server string, e Evaluation) (string, int) {
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

	return string(body), resp.StatusCode
}
