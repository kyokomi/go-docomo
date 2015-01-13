package docomo

import (
	"testing"
)

func TestDialogueGet(t *testing.T) {

	type TestCase struct {
		in  string
		out string
	}

	testCase := TestCase{
		in:  "tests/stubs/dialogue.json",
		out: "こんにちは。ちょうど退屈してたんだ。",
	}

	serve, client := Stub(testCase.in)
	defer serve.Close()

	d := DialogueRequest{}
	res, err := client.Dialogue.Get(d, true)
	if err != nil {
		t.Errorf("error Request %s\n", err)
	}

	if res.Utt != testCase.out {
		t.Errorf("error Response %s != \n", res.Utt, testCase.out)
	}
}
