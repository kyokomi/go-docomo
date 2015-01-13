package docomo

import "testing"

func TestKnowledgeGet(t *testing.T) {

	type TestCase struct {
		in  string
		out string
	}

	testCase := TestCase{
		in:  "../tests/stubs/knowledge.json",
		out: "ガガーリン",
	}

	serve, client := Stub(testCase.in)
	defer serve.Close()

	req := KnowledgeQARequest{}
	res, err := client.KnowledgeQA.Get(req)
	if err != nil {
		t.Errorf("error Request %s\n", err)
	}

	answers := 5
	if len(res.Answers) != answers {
		t.Errorf("error Response answers lenght %d != %d\n", len(res.Answers), answers)
	}

	if res.Answers[0].AnswerText != testCase.out {
		t.Errorf("error Response %s != %s\n", res.Answers[0].AnswerText, testCase.out)
	}
}
