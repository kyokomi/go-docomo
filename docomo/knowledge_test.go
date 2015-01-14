package docomo

import (
	"testing"
	"reflect"
)

func TestKnowledgeGet(t *testing.T) {

	type TestCase struct {
		in  string
		out KnowledgeQAResponse
	}

	testCase := TestCase{
		in:  "../tests/stubs/knowledge.json",
	}

	serve, client := Stub(testCase.in, &testCase.out)
	defer serve.Close()

	req := KnowledgeQARequest{}
	res, err := client.KnowledgeQA.Get(req)
	if err != nil {
		t.Errorf("error Request %s\n", err)
	}

	if !reflect.DeepEqual(*res, testCase.out) {
		t.Errorf("error Response %s != %s\n", res, testCase.out)
	}
}
