package docomo

import (
	"testing"
	"reflect"
)

func TestDialogueGet(t *testing.T) {

	type TestCase struct {
		in  string
		out DialogueResponse
	}

	testCase := TestCase{
		in:  "../tests/stubs/dialogue.json",
	}
	serve, client := Stub(testCase.in, &testCase.out)
	defer serve.Close()

	d := DialogueRequest{}
	res, err := client.Dialogue.Get(d, true)
	if err != nil {
		t.Errorf("error Request %s\n", err)
	}

	if !reflect.DeepEqual(*res, testCase.out) {
		t.Errorf("error Response %s != %s\n", res, testCase.out)
	}
}
