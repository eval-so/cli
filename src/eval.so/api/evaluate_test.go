package evaluate

import (
	"."
	"strings"
	"testing"
)

func TestEvaluate(t *testing.T) {
	eval := evaluate.Evaluation{
		Code:     "puts 123",
		Language: "ruby",
	}

	// These aren't meant to test the API itself, just that we use it properly.

	result, statusCode := evaluate.Evaluate("http://eval.so/api/evaluate", eval)

	if statusCode != 200 {
		t.Error("Did not get a HTTP 200 status back")
	}

	if !strings.HasPrefix(result.Stdout, "123") {
		t.Error("STDOUT did not start with '123'")
	}
}
